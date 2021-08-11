package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTrx{}

func NewMsgCreateTrx(creator string, index string, chainID int32, address string, crossChainMsgs string, proof string) *MsgCreateTrx {
	return &MsgCreateTrx{
		Creator:        creator,
		Index:          index,
		ChainID:        chainID,
		Address:        address,
		CrossChainMsgs: crossChainMsgs,
		Proof:          proof,
	}
}

func (msg *MsgCreateTrx) Route() string {
	return RouterKey
}

func (msg *MsgCreateTrx) Type() string {
	return "CreateTrx"
}

func (msg *MsgCreateTrx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTrx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTrx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTrx{}

func NewMsgUpdateTrx(creator string, index string, chainID int32, address string, crossChainMsgs string) *MsgUpdateTrx {
	return &MsgUpdateTrx{
		Creator:        creator,
		Index:          index,
		ChainID:        chainID,
		Address:        address,
		CrossChainMsgs: crossChainMsgs,
	}
}

func (msg *MsgUpdateTrx) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTrx) Type() string {
	return "UpdateTrx"
}

func (msg *MsgUpdateTrx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTrx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTrx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTrx{}

func NewMsgDeleteTrx(creator string, index string) *MsgDeleteTrx {
	return &MsgDeleteTrx{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteTrx) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTrx) Type() string {
	return "DeleteTrx"
}

func (msg *MsgDeleteTrx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTrx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTrx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
