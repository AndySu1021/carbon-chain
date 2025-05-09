package simulation

import (
	"math/rand"

	"carbonchain/x/voting/keeper"
	"carbonchain/x/voting/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgVote(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgVote{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Vote simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "Vote simulation not implemented"), nil, nil
	}
}
