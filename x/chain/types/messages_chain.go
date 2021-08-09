package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateChain{}

func NewMsgCreateChain(creator string, index string, sourceChainID int32, height int32, proof string, proxyAddress string, extra string, headerOrCrossChainMsg string) *MsgCreateChain {
	return &MsgCreateChain{
		Creator:               creator,
		Index:                 index,
		SourceChainID:         sourceChainID,
		Height:                height,
		Proof:                 proof,
		ProxyAddress:          proxyAddress,
		Extra:                 extra,
		HeaderOrCrossChainMsg: headerOrCrossChainMsg,
	}
}

func (msg *MsgCreateChain) Route() string {
	return RouterKey
}

func (msg *MsgCreateChain) Type() string {
	return "CreateChain"
}

func (msg *MsgCreateChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateChain{}

func NewMsgUpdateChain(creator string, index string, sourceChainID int32, height int32, proof string, proxyAddress string, extra string, headerOrCrossChainMsg string) *MsgUpdateChain {
	return &MsgUpdateChain{
		Creator:               creator,
		Index:                 index,
		SourceChainID:         sourceChainID,
		Height:                height,
		Proof:                 proof,
		ProxyAddress:          proxyAddress,
		Extra:                 extra,
		HeaderOrCrossChainMsg: headerOrCrossChainMsg,
	}
}

func (msg *MsgUpdateChain) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChain) Type() string {
	return "UpdateChain"
}

func (msg *MsgUpdateChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteChain{}

func NewMsgDeleteChain(creator string, index string) *MsgDeleteChain {
	return &MsgDeleteChain{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteChain) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChain) Type() string {
	return "DeleteChain"
}

func (msg *MsgDeleteChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
