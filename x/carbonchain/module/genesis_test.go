package carbonchain_test

import (
	"testing"

	keepertest "carbonchain/testutil/keeper"
	"carbonchain/testutil/nullify"
	carbonchain "carbonchain/x/carbonchain/module"
	"carbonchain/x/carbonchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarbonchainKeeper(t)
	carbonchain.InitGenesis(ctx, k, genesisState)
	got := carbonchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
