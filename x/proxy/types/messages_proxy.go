package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProxy{}

func NewMsgCreateProxy(creator string, index string, addressList string, address string) *MsgCreateProxy {
	return &MsgCreateProxy{
		Creator:     creator,
		Index:       index,
		AddressList: addressList,
		Address:     address,
	}
}

func (msg *MsgCreateProxy) Route() string {
	return RouterKey
}

func (msg *MsgCreateProxy) Type() string {
	return "CreateProxy"
}

func (msg *MsgCreateProxy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProxy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProxy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProxy{}

func NewMsgUpdateProxy(creator string, index string, addressList string, address string) *MsgUpdateProxy {
	return &MsgUpdateProxy{
		Creator:     creator,
		Index:       index,
		AddressList: addressList,
		Address:     address,
	}
}

func (msg *MsgUpdateProxy) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProxy) Type() string {
	return "UpdateProxy"
}

func (msg *MsgUpdateProxy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProxy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProxy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProxy{}

func NewMsgDeleteProxy(creator string, index string) *MsgDeleteProxy {
	return &MsgDeleteProxy{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteProxy) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProxy) Type() string {
	return "DeleteProxy"
}

func (msg *MsgDeleteProxy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProxy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProxy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
