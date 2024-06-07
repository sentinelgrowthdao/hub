// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/sentinel-official/hub/v12/x/provider/types/v1"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1types.RegisterInterfaces(registry)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterRequest{},
		&MsgUpdateRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
