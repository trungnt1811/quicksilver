package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ingenuity-build/quicksilver/x/participationrewards/types"
)

// HandleAddProtocolDataProposal is a handler for executing a passed add protocol data proposal
func HandleAddProtocolDataProposal(ctx sdk.Context, k Keeper, p *types.AddProtocolDataProposal) error {
	protocolData := NewProtocolData(p.Type, p.Protocol, p.Data)

	pdtv, exists := types.ProtocolDataType_value[p.Type]
	if !exists {
		return types.ErrUnknownProtocolDataType
	}

	_, err := types.UnmarshalProtocolData(pdtv, p.Data)
	if err != nil {
		return err
	}

	k.SetProtocolData(ctx, p.Key, protocolData)

	return nil
}
