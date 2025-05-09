package keeper

import (
	"carbonchain/x/voting/types"
	"context"
)

func (k msgServer) CreateProposal(goCtx context.Context, msg *types.MsgCreateProposal) (*types.MsgCreateProposalResponse, error) {
	id, err := k.AppendProposal(goCtx, types.Proposal{
		Title:       msg.Title,
		Description: msg.Description,
		Creator:     msg.Creator,
	})

	if err != nil {
		return nil, err
	}

	//sdkCtx := sdk.UnwrapSDKContext(goCtx)
	//sdkCtx.EventManager().EmitEvents(sdk.Events{
	//	sdk.NewEvent(
	//		types.EventTypeTransfer,
	//		sdk.NewAttribute(types.AttributeKeyRecipient, toAddr.String()),
	//		sdk.NewAttribute(types.AttributeKeySender, fromAddrString),
	//		sdk.NewAttribute(sdk.AttributeKeyAmount, amt.String()),
	//	),
	//	sdk.NewEvent(
	//		sdk.EventTypeMessage,
	//		sdk.NewAttribute(types.AttributeKeySender, fromAddrString),
	//	),
	//})

	return &types.MsgCreateProposalResponse{Id: id}, nil
}
