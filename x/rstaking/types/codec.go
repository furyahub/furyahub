package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetInflationBase{}, "rstaking/SetInflationBase", nil)
	cdc.RegisterConcrete(&MsgAddValToWhitelist{}, "rstaking/AddValToWhitelist", nil)
	cdc.RegisterConcrete(&MsgToggleWhitelistSwitch{}, "rstaking/ToggleWhitelistSwitch", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "rstaking/Withdraw", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetInflationBase{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddValToWhitelist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleWhitelistSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
