package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCrossChainTransfer{}

func NewMsgCrossChainTransfer(creator string, asset string, amount int32, fromChain string, fromAddress string, toChain string, toAddress string) *MsgCrossChainTransfer {
	return &MsgCrossChainTransfer{
		Creator:     creator,
		Asset:       asset,
		Amount:      amount,
		FromChain:   fromChain,
		FromAddress: fromAddress,
		ToChain:     toChain,
		ToAddress:   toAddress,
	}
}

func (msg *MsgCrossChainTransfer) Route() string {
	return RouterKey
}

func (msg *MsgCrossChainTransfer) Type() string {
	return "CrossChainTransfer"
}

func (msg *MsgCrossChainTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCrossChainTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrossChainTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
