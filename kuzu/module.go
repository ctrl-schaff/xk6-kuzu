package kuzu

import (
	"context"

	"github.com/kuzudb/go-kuzu"
	"go.k6.io/k6/js/modules"
)

// ImportPath contains module's JavaScript import path.
const ImportPath = "k6/x/kuzu"

// New creates a new instance of the extension's JavaScript module.
func New() modules.Module {
	return new(rootModule)
}

// rootModule is the global module object type. It is instantiated once per test
// run and will be used to create `k6/x/kuzu` module instances for each VU.
type rootModule struct{}

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*rootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	instance := &module{}

	instance.exports.Default = instance
	instance.exports.Named = map[string]any{
		"open": instance.OpenConnection,
	}

	instance.vu = vu

	return instance
}

// module represents an instance of the JavaScript module for every VU.
type module struct {
	exports modules.Exports
	vu      modules.VU
}

// Exports is representation of ESM exports of a module.
func (mod *module) Exports() modules.Exports {
	return mod.exports
}

func (kmod *module) OpenConnection(databasePath string, opts *kuzuOptions) (*DatabaseConnection, error) {
	kuzuConfig := opts.apply()
	db, err := kuzu.OpenDatabase(databasePath, kuzuConfig)
	if err != nil {
		return nil, err
	}

	conn, err := kuzu.OpenConnection(db)
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{conn: conn, ctx: kmod.vu.Context}, nil
}

// Database is a database handle representing a pool of zero or more underlying connections.
type DatabaseConnection struct {
	conn *kuzu.Connection
	ctx  func() context.Context
}

func (kuzuConn *DatabaseConnection) Query(queryString string) (*kuzu.QueryResult, error) {
	queryResult, err := kuzuConn.conn.Query(queryString)
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}

// Close the database and prevents new queries from starting.
func (kuzuConn *DatabaseConnection) Close() {
	kuzuConn.conn.Close()
}
