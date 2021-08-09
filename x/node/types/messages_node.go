package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateNode{}

func NewMsgCreateNode(creator string, index string, peerPubkey string, address string) *MsgCreateNode {
	return &MsgCreateNode{
		Creator:    creator,
		Index:      index,
		PeerPubkey: peerPubkey,
		Address:    address,
	}
}

func (msg *MsgCreateNode) Route() string {
	return RouterKey
}

func (msg *MsgCreateNode) Type() string {
	return "CreateNode"
}

func (msg *MsgCreateNode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNode{}

func NewMsgUpdateNode(creator string, index string, peerPubkey string, address string) *MsgUpdateNode {
	return &MsgUpdateNode{
		Creator:    creator,
		Index:      index,
		PeerPubkey: peerPubkey,
		Address:    address,
	}
}

func (msg *MsgUpdateNode) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNode) Type() string {
	return "UpdateNode"
}

func (msg *MsgUpdateNode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNode{}

func NewMsgDeleteNode(creator string, index string) *MsgDeleteNode {
	return &MsgDeleteNode{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteNode) Route() string {
	return RouterKey
}

func (msg *MsgDeleteNode) Type() string {
	return "DeleteNode"
}

func (msg *MsgDeleteNode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
