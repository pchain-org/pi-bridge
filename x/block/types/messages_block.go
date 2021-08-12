package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBlock{}

func NewMsgCreateBlock(creator string, index string, chainID int32, address string, headers []string) *MsgCreateBlock {
	return &MsgCreateBlock{
		Creator: creator,
		Index:   index,
		ChainID: chainID,
		Address: address,
		Headers: headers,
	}
}

func (msg *MsgCreateBlock) Route() string {
	return RouterKey
}

func (msg *MsgCreateBlock) Type() string {
	return "CreateBlock"
}

func (msg *MsgCreateBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBlock{}

func NewMsgUpdateBlock(creator string, index string, chainID int32, address string, headers []string) *MsgUpdateBlock {
	return &MsgUpdateBlock{
		Creator: creator,
		Index:   index,
		ChainID: chainID,
		Address: address,
		Headers: headers,
	}
}

func (msg *MsgUpdateBlock) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBlock) Type() string {
	return "UpdateBlock"
}

func (msg *MsgUpdateBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBlock{}

func NewMsgDeleteBlock(creator string, index string) *MsgDeleteBlock {
	return &MsgDeleteBlock{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteBlock) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBlock) Type() string {
	return "DeleteBlock"
}

func (msg *MsgDeleteBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
