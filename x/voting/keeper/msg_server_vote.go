package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"

	"carbonchain/x/voting/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal, err := k.GetProposal(ctx, msg.ProposalId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPanic, "proposal %d not found", msg.ProposalId)
	}

	if msg.Vote == "yes" {
		proposal.YesCount++
	} else if msg.Vote == "no" {
		proposal.NoCount++
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "invalid vote option")
	}

	store := k.storeService.OpenKVStore(goCtx)
	bz := k.cdc.MustMarshal(&proposal)
	if err = store.Set(sdk.Uint64ToBigEndian(msg.ProposalId), bz); err != nil {
		return nil, err
	}

	return &types.MsgVoteResponse{}, nil
}
