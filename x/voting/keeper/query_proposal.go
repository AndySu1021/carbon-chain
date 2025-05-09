package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"

	"carbonchain/x/voting/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Proposal(goCtx context.Context, req *types.QueryProposalRequest) (*types.QueryProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal, err := k.GetProposal(ctx, req.Id)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPanic, "proposal %d not found", req.Id)
	}

	return &types.QueryProposalResponse{Proposal: &proposal}, nil
}
