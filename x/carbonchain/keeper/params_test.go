package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "carbonchain/testutil/keeper"
	"carbonchain/x/carbonchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CarbonchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
