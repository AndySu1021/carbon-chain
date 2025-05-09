package keeper

import (
	"carbonchain/x/voting/types"
)

var _ types.QueryServer = Keeper{}
