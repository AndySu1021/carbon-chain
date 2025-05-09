package keeper

import (
	"carbonchain/x/voting/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendProposal(ctx context.Context, p types.Proposal) (uint64, error) {
	store := k.storeService.OpenKVStore(ctx)
	countBytes, err := store.Get([]byte("proposal-count"))
	if err != nil {
		return 0, nil
	}

	count := sdk.BigEndianToUint64(countBytes)

	p.Id = count
	bz := k.cdc.MustMarshal(&p)
	if err = store.Set(sdk.Uint64ToBigEndian(count), bz); err != nil {
		return 0, err
	}

	// set proposal count
	if err = store.Set([]byte("proposal-count"), sdk.Uint64ToBigEndian(count+1)); err != nil {
		return 0, err
	}

	return p.Id, nil
}

func (k Keeper) GetProposal(ctx context.Context, id uint64) (types.Proposal, error) {
	store := k.storeService.OpenKVStore(ctx)
	proposalBytes, err := store.Get(sdk.Uint64ToBigEndian(id))
	if err != nil {
		return types.Proposal{}, err
	}

	var proposal types.Proposal
	k.cdc.MustUnmarshal(proposalBytes, &proposal)

	return proposal, nil
}
