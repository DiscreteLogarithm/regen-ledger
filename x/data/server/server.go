package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/store/lookup"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
)

type serverImpl struct {
	storeKey   sdk.StoreKey
	iriIdTable lookup.Table
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	tbl, err := lookup.NewTable([]byte{IriIdTablePrefix})
	if err != nil {
		panic(err)
	}

	return serverImpl{
		storeKey:   storeKey,
		iriIdTable: tbl,
	}
}

func RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
