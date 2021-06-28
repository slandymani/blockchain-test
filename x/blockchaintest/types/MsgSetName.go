package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetName{}

type MsgSetName struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Value string         `json:"value" yaml:"value"`
	Name  string         `json:"name" yaml:"name"`
}

func NewMsgSetName(creator sdk.AccAddress, value string, name string) MsgSetName {
	return MsgSetName{
		Owner: creator,
		Value: value,
		Name:  name,
	}
}

func (msg MsgSetName) Route() string {
	return RouterKey
}

func (msg MsgSetName) Type() string {
	return "set_name"
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}
