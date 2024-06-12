// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	deposittypes "github.com/sentinel-official/hub/v12/x/deposit/types"
	v1nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v1"
	v2nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	v1plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v1"
	v2plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v2"
	v1providertypes "github.com/sentinel-official/hub/v12/x/provider/types/v1"
	v2providertypes "github.com/sentinel-official/hub/v12/x/provider/types/v2"
	v1sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v1"
	v2sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v2"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	deposittypes.RegisterInterfaces(registry)
	subscriptiontypes.RegisterInterfaces(registry)
	v1nodetypes.RegisterInterfaces(registry)
	v1plantypes.RegisterInterfaces(registry)
	v1providertypes.RegisterInterfaces(registry)
	v1sessiontypes.RegisterInterfaces(registry)
	v2nodetypes.RegisterInterfaces(registry)
	v2plantypes.RegisterInterfaces(registry)
	v2providertypes.RegisterInterfaces(registry)
	v2sessiontypes.RegisterInterfaces(registry)
}
