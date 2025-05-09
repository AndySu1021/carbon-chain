package voting

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"carbonchain/testutil/sample"
	votingsimulation "carbonchain/x/voting/simulation"
	"carbonchain/x/voting/types"
)

// avoid unused import issue
var (
	_ = votingsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateProposal = "op_weight_msg_create_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProposal int = 100

	opWeightMsgVote = "op_weight_msg_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	votingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&votingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = defaultWeightMsgCreateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProposal,
		votingsimulation.SimulateMsgCreateProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVote int
	simState.AppParams.GetOrGenerate(opWeightMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = defaultWeightMsgVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVote,
		votingsimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProposal,
			defaultWeightMsgCreateProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				votingsimulation.SimulateMsgCreateProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVote,
			defaultWeightMsgVote,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				votingsimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
