package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCrossChainTransfer{}, "trx/CrossChainTransfer", nil)

	cdc.RegisterConcrete(&MsgCreateTrx{}, "trx/CreateTrx", nil)
	cdc.RegisterConcrete(&MsgUpdateTrx{}, "trx/UpdateTrx", nil)
	cdc.RegisterConcrete(&MsgDeleteTrx{}, "trx/DeleteTrx", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrossChainTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTrx{},
		&MsgUpdateTrx{},
		&MsgDeleteTrx{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
