package voting

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "carbonchain/api/carbonchain/voting"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Proposal",
					Use:            "proposal [id]",
					Short:          "Query proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateProposal",
					Use:            "create-proposal [title] [description]",
					Short:          "Send a create-proposal tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "Vote",
					Use:            "vote [proposal-id] [vote]",
					Short:          "Send a vote tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proposalId"}, {ProtoField: "vote"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
