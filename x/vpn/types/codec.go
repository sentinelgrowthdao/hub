// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	deposittypes "github.com/sentinel-official/hub/v12/x/deposit/types"
	v1nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v1"
	v2nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	v1plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v1"
	v2plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v2"
	providertypes "github.com/sentinel-official/hub/v12/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	deposittypes.RegisterInterfaces(registry)
	providertypes.RegisterInterfaces(registry)
	sessiontypes.RegisterInterfaces(registry)
	subscriptiontypes.RegisterInterfaces(registry)
	v1nodetypes.RegisterInterfaces(registry)
	v1plantypes.RegisterInterfaces(registry)
	v2nodetypes.RegisterInterfaces(registry)
	v2plantypes.RegisterInterfaces(registry)
}
