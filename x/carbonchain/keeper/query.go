package keeper

import (
	"carbonchain/x/carbonchain/types"
)

var _ types.QueryServer = Keeper{}
