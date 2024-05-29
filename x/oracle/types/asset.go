package types

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/v12/third_party/osmosis/x/poolmanager/client/queryproto"
	protorevtypes "github.com/sentinel-official/hub/v12/third_party/osmosis/x/protorev/types"
)

func (a *Asset) QueryGetProtoRevPoolRequest(cdc codec.Codec) abcitypes.RequestQuery {
	return abcitypes.RequestQuery{
		Data: cdc.MustMarshal(
			&protorevtypes.QueryGetProtoRevPoolRequest{
				BaseDenom:  a.BaseAssetDenom,
				OtherDenom: a.QuoteAssetDenom,
			},
		),
		Path: "/osmosis.protorev.v1beta1.Query/GetProtoRevPool",
	}
}

func (a *Asset) SpotPriceRequest(cdc codec.Codec) abcitypes.RequestQuery {
	return abcitypes.RequestQuery{
		Data: cdc.MustMarshal(
			&queryproto.SpotPriceRequest{
				PoolId:          a.PoolID,
				BaseAssetDenom:  a.BaseAssetDenom,
				QuoteAssetDenom: a.QuoteAssetDenom,
			},
		),
		Path: "/osmosis.poolmanager.v1beta1.Query/SpotPrice",
	}
}
