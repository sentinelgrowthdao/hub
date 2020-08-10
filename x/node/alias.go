// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/node/types
// ALIASGEN: github.com/sentinel-official/hub/x/node/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/node/querier
package node

import (
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/querier"
	"github.com/sentinel-official/hub/x/node/types"
)

const (
	Codespace               = types.Codespace
	EventTypeSet            = types.EventTypeSet
	EventTypeUpdate         = types.EventTypeUpdate
	EventTypeSetStatus      = types.EventTypeSetStatus
	AttributeKeyProvider    = types.AttributeKeyProvider
	AttributeKeyAddress     = types.AttributeKeyAddress
	AttributeKeyStatus      = types.AttributeKeyStatus
	ModuleName              = types.ModuleName
	ParamsSubspace          = types.ParamsSubspace
	QuerierRoute            = types.QuerierRoute
	CategoryUnknown         = types.CategoryUnknown
	DefaultInactiveDuration = types.DefaultInactiveDuration
	QueryNode               = types.QueryNode
	QueryNodes              = types.QueryNodes
	QueryNodesForProvider   = types.QueryNodesForProvider
)

var (
	// functions aliases
	RegisterCodec                  = types.RegisterCodec
	ErrorMarshal                   = types.ErrorMarshal
	ErrorUnmarshal                 = types.ErrorUnmarshal
	ErrorUnknownMsgType            = types.ErrorUnknownMsgType
	ErrorUnknownQueryType          = types.ErrorUnknownQueryType
	ErrorInvalidField              = types.ErrorInvalidField
	ErrorProviderDoesNotExist      = types.ErrorProviderDoesNotExist
	ErrorDuplicateNode             = types.ErrorDuplicateNode
	ErrorNodeDoesNotExist          = types.ErrorNodeDoesNotExist
	ErrorCanNotUpdate              = types.ErrorCanNotUpdate
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	NodeKey                        = types.NodeKey
	GetNodeForProviderKeyPrefix    = types.GetNodeForProviderKeyPrefix
	NodeForProviderKey             = types.NodeForProviderKey
	GetActiveNodeAtKeyPrefix       = types.GetActiveNodeAtKeyPrefix
	ActiveNodeAtKey                = types.ActiveNodeAtKey
	NewMsgRegister                 = types.NewMsgRegister
	NewMsgUpdate                   = types.NewMsgUpdate
	NewMsgSetStatus                = types.NewMsgSetStatus
	CategoryFromString             = types.CategoryFromString
	NewParams                      = types.NewParams
	DefaultParams                  = types.DefaultParams
	ParamsKeyTable                 = types.ParamsKeyTable
	NewQueryNodeParams             = types.NewQueryNodeParams
	NewQueryNodesParams            = types.NewQueryNodesParams
	NewQueryNodesForProviderParams = types.NewQueryNodesForProviderParams
	NewKeeper                      = keeper.NewKeeper
	Querier                        = querier.Querier

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	EventModuleName          = types.EventModuleName
	NodeKeyPrefix            = types.NodeKeyPrefix
	NodeForProviderKeyPrefix = types.NodeForProviderKeyPrefix
	ActiveNodeAtKeyPrefix    = types.ActiveNodeAtKeyPrefix
	KeyInactiveDuration      = types.KeyInactiveDuration
)

type (
	GenesisState                = types.GenesisState
	MsgRegister                 = types.MsgRegister
	MsgUpdate                   = types.MsgUpdate
	MsgSetStatus                = types.MsgSetStatus
	Category                    = types.Category
	Node                        = types.Node
	Nodes                       = types.Nodes
	Params                      = types.Params
	QueryNodeParams             = types.QueryNodeParams
	QueryNodesParams            = types.QueryNodesParams
	QueryNodesForProviderParams = types.QueryNodesForProviderParams
	Keeper                      = keeper.Keeper
)
