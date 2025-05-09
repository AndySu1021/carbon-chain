package voting_test

import (
	"testing"

	keepertest "carbonchain/testutil/keeper"
	"carbonchain/testutil/nullify"
	voting "carbonchain/x/voting/module"
	"carbonchain/x/voting/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VotingKeeper(t)
	voting.InitGenesis(ctx, k, genesisState)
	got := voting.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
