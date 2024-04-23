package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name.
	ModuleName = "participationrewards"
	// StoreKey defines the primary module store key.
	StoreKey = ModuleName
	// QuerierRoute is the querier route for the participationrewards store.
	QuerierRoute = StoreKey
	// RouterKey is the message route for participationrewards.
	RouterKey = ModuleName

	OsmosisParamsKey  = "osmosisparams"
	UmeeParamsKey     = "umeeparams"
	CrescentParamsKey = "crescentparams"
	ProofTypeBank     = "bank"
	ProofTypeLeverage = "leverage"
	ProofTypeLPFarm   = "lpfarm"
)

var (
	KeyPrefixProtocolData        = []byte{0x00}
	KeyPrefixHoldingAllocation   = []byte{0x01}
	KeyPrefixValidatorAllocation = []byte{0x02}
	KeyPrefixValues              = []byte{0x03}
)

func GetProtocolDataKey(pdType ProtocolDataType, key []byte) []byte {
	return append(sdk.Uint64ToBigEndian(uint64(pdType)), key...)
}

func GetPrefixProtocolDataKey(pdType ProtocolDataType) []byte {
	return sdk.Uint64ToBigEndian(uint64(pdType))
}
