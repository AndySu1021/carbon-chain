package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProposal{}

func NewMsgCreateProposal(creator string, title string, description string) *MsgCreateProposal {
	return &MsgCreateProposal{
		Creator:     creator,
		Title:       title,
		Description: description,
	}
}

func (msg *MsgCreateProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
