package keeper

import (
	"errors"
	"fmt"

	"github.com/ingenuity-build/multierror"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/quicksilver-zone/quicksilver/utils"
	"github.com/quicksilver-zone/quicksilver/utils/addressutils"
	icstypes "github.com/quicksilver-zone/quicksilver/x/interchainstaking/types"
	"github.com/quicksilver-zone/quicksilver/x/participationrewards/types"
)

type TokenValues map[string]sdk.Dec

type (
	AssetGraph      map[string]map[string]sdk.Dec
	AssetGraphSlice map[string]map[string][]sdk.Dec
)

func DepthFirstSearch(graph AssetGraph, visited map[string]struct{}, asset string, price sdk.Dec, result TokenValues) {
	visited[asset] = struct{}{}
	result[asset] = price

	for _, neighbour := range utils.Keys(graph[asset]) {
		if _, ok := visited[neighbour]; !ok {
			DepthFirstSearch(graph, visited, neighbour, graph[asset][neighbour].Mul(price), result)
		}
	}
}

func (k *Keeper) CalcTokenValues(ctx sdk.Context) (TokenValues, error) {
	k.Logger(ctx).Info("calcTokenValues")

	data, found := k.GetProtocolData(ctx, types.ProtocolDataTypeOsmosisParams, "osmosisparams")
	if !found {
		return TokenValues{}, errors.New("could not find osmosisparams protocol data")
	}
	osmoParams, err := types.UnmarshalProtocolData(types.ProtocolDataTypeOsmosisParams, data.Data)
	if err != nil {
		return TokenValues{}, err
	}

	baseDenom := osmoParams.(*types.OsmosisParamsProtocolData).BaseDenom

	tvs := make(TokenValues)
	graph := make(AssetGraphSlice)
	graph2 := make(AssetGraph)

	// add base value
	tvs[baseDenom] = sdk.OneDec()

	// capture errors from iterator
	errs := make(map[string]error)
	k.IteratePrefixedProtocolDatas(ctx, types.GetPrefixProtocolDataKey(types.ProtocolDataTypeOsmosisPool), func(idx int64, _ []byte, data types.ProtocolData) bool {
		idxLabel := fmt.Sprintf("index[%d]", idx)
		ipool, err := types.UnmarshalProtocolData(types.ProtocolDataTypeOsmosisPool, data.Data)
		if err != nil {
			errs[idxLabel] = err
			return true
		}
		pool, _ := ipool.(*types.OsmosisPoolProtocolData)

		// pool must be a base pair
		if len(pool.Denoms) != 2 {
			return false
		}

		if pool.PoolData == nil {
			errs[idxLabel] = fmt.Errorf("pool data is nil, awaiting OsmosisPoolUpdateCallback")
			return true
		}
		gammPool, err := pool.GetPool()
		if err != nil {
			errs[idxLabel] = err
			return true
		}

		denoms := utils.Keys(pool.Denoms)

		prettyDenom0 := pool.Denoms[denoms[0]].Denom
		prettyDenom1 := pool.Denoms[denoms[1]].Denom

		for _, ibcDenom := range utils.Keys(pool.Denoms) {
			if _, ok := graph[pool.Denoms[ibcDenom].Denom]; !ok {
				graph[pool.Denoms[ibcDenom].Denom] = make(map[string][]sdk.Dec)
			}
		}

		value, err := gammPool.SpotPrice(ctx, denoms[0], denoms[1])
		if err != nil {
			errs[idxLabel] = err
			return true
		}

		decVal := sdk.NewDecFromBigIntWithPrec(value.Dec().BigInt(), 18)

		graph[prettyDenom0][prettyDenom1] = append(graph[prettyDenom0][prettyDenom1], decVal)
		graph[prettyDenom1][prettyDenom0] = append(graph[prettyDenom1][prettyDenom0], sdk.OneDec().Quo(decVal))

		return false
	})

	k.IteratePrefixedProtocolDatas(ctx, types.GetPrefixProtocolDataKey(types.ProtocolDataTypeOsmosisCLPool), func(idx int64, _ []byte, data types.ProtocolData) bool {
		idxLabel := fmt.Sprintf("index[%d]", idx)
		ipool, err := types.UnmarshalProtocolData(types.ProtocolDataTypeOsmosisCLPool, data.Data)
		if err != nil {
			errs[idxLabel] = err
			return true
		}
		pool, _ := ipool.(*types.OsmosisClPoolProtocolData)

		// pool must be a base pair
		if len(pool.Denoms) != 2 {
			return false
		}

		if pool.PoolData == nil {
			errs[idxLabel] = fmt.Errorf("pool data is nil, awaiting OsmosisClPoolUpdateCallback")
			return true
		}
		clPool, err := pool.GetPool()
		if err != nil {
			errs[idxLabel] = err
			return true
		}

		denoms := utils.Keys(pool.Denoms)
		prettyDenom0 := pool.Denoms[denoms[0]].Denom
		prettyDenom1 := pool.Denoms[denoms[1]].Denom

		for _, ibcDenom := range utils.Keys(pool.Denoms) {
			if _, ok := graph[pool.Denoms[ibcDenom].Denom]; !ok {
				graph[pool.Denoms[ibcDenom].Denom] = make(map[string][]sdk.Dec)
			}
		}

		value, err := clPool.SpotPrice(ctx, denoms[0], denoms[1])
		if err != nil {
			errs[idxLabel] = err
			return true
		}

		decVal := sdk.NewDecFromBigIntWithPrec(value.Dec().BigInt(), 18)

		graph[prettyDenom0][prettyDenom1] = append(graph[prettyDenom0][prettyDenom1], decVal)
		graph[prettyDenom1][prettyDenom0] = append(graph[prettyDenom1][prettyDenom0], sdk.OneDec().Quo(decVal))

		return false
	})

	for _, denom0 := range utils.Keys(graph) {
		graph2[denom0] = make(map[string]sdk.Dec)
		values := graph[denom0]
		for _, denom1 := range utils.Keys(values) {
			value := sdk.ZeroDec()
			count := math.ZeroInt()
			for _, asset := range values[denom1] {
				value = value.Add(asset)
				count = count.Add(math.OneInt())
			}
			graph2[denom0][denom1] = value.QuoInt(count)
		}
	}

	visited := make(map[string]struct{})
	DepthFirstSearch(graph2, visited, baseDenom, sdk.OneDec(), tvs)

	if len(errs) > 0 {
		return nil, multierror.New(errs)
	}

	return tvs, nil
}

// SetZoneAllocations returns the proportional zone rewards allocations as a
// map indexed by the zone id.
func (k *Keeper) SetZoneAllocations(ctx sdk.Context, tvs TokenValues) error {
	holdingAllocation := k.GetHoldingAllocation(ctx, types.ModuleName)
	validatorAllocation := k.GetValidatorAllocation(ctx, types.ModuleName)
	k.Logger(ctx).Info("setZoneAllocations", "holdingAllocation", holdingAllocation, "validatorAllocation", validatorAllocation)

	otvl := sdk.ZeroDec()
	// pass 1: iterate zones - set tvl & calc overall tvl
	k.icsKeeper.IterateZones(ctx, func(index int64, zone *icstypes.Zone) (stop bool) {
		tv, exists := tvs[zone.BaseDenom]
		if !exists {
			k.Logger(ctx).Error(fmt.Sprintf("unable to obtain token value for zone %s", zone.ChainId))
			return false
		}
		ztvl := sdk.NewDecFromInt(k.icsKeeper.GetDelegatedAmount(ctx, zone).Amount.Add(k.icsKeeper.GetDelegationsInProcess(ctx, zone.ChainId))).Mul(tv)
		zone.Tvl = ztvl
		k.icsKeeper.SetZone(ctx, zone)

		k.Logger(ctx).Info("zone tvl", "zone", zone.ChainId, "tvl", ztvl)

		otvl = otvl.Add(ztvl)
		return false
	})

	// check overall protocol tvl
	if otvl.IsZero() {
		err := errors.New("protocol tvl is zero")
		return err
	}

	// pass 2: iterate zones - calc zone tvl proportion & set allocations
	k.icsKeeper.IterateZones(ctx, func(index int64, zone *icstypes.Zone) (stop bool) {
		if zone.Tvl.IsNil() {
			return false
		}

		zp := zone.Tvl.Quo(otvl)
		k.Logger(ctx).Info("zone proportion", "zone", zone.ChainId, "proportion", zp)
		k.SetValidatorAllocation(ctx, zone.ChainId, sdk.NewCoin(validatorAllocation.Denom, sdk.NewDecFromInt(validatorAllocation.Amount).Mul(zp).TruncateInt()))
		k.SetHoldingAllocation(ctx, zone.ChainId, sdk.NewCoin(holdingAllocation.Denom, sdk.NewDecFromInt(holdingAllocation.Amount).Mul(zp).TruncateInt()))
		return false
	})

	return nil
}

// DistributeToUsersFromModule sends the allocated user rewards to the user address.
func (k *Keeper) DistributeToUsersFromModule(ctx sdk.Context, userAllocations []types.UserAllocation) error {
	k.Logger(ctx).Info("distribute to users from module", "allocations", userAllocations)

	for _, ua := range userAllocations {
		if !ua.Amount.IsPositive() {
			continue
		}

		coins := sdk.NewCoins(ua.Amount)

		addrBytes, err := addressutils.AccAddressFromBech32(ua.Address, "")
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addrBytes, coins)
		if err != nil {
			return err
		}
		k.Logger(ctx).Info("distribute to user", "address", ua.Address, "coins", coins, "remaining", k.GetModuleBalance(ctx))

	}

	return nil
}

// DistributeToUsers sends the allocated user rewards to the user address.
func (k *Keeper) DistributeToUsersFromAddress(ctx sdk.Context, userAllocations []types.UserAllocation, fromAddress string) error {
	k.Logger(ctx).Info("distribute to users from account", "allocations", userAllocations)

	fromAddrBytes, err := addressutils.AccAddressFromBech32(fromAddress, "")
	if err != nil {
		return err
	}

	for _, ua := range userAllocations {
		if !ua.Amount.IsPositive() {
			continue
		}

		coins := sdk.NewCoins(
			ua.Amount,
		)

		addrBytes, err := addressutils.AccAddressFromBech32(ua.Address, "")
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoins(ctx, fromAddrBytes, addrBytes, coins)
		if err != nil {
			return err
		}
		k.Logger(ctx).Info("distribute to user", "address", ua.Address, "coins", coins, "remaining", k.GetModuleBalance(ctx))
	}

	return nil
}
