package kuzu

import (
	"context"

	"github.com/grafana/sobek"
	"github.com/kuzudb/go-kuzu"
	"go.k6.io/k6/js/modules"
)

type RootModule struct{}
type KuzuModule struct{ vu modules.VU }

func NewModule() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &KuzuModule{vu}
}

func (km *KuzuModule) Exports() modules.Exports {
	return modules.Exports{Default: km}
}

var _ modules.Module = (*RootModule)(nil)
var _ modules.Instance = (*KuzuModule)(nil)

// options represents connection related options for the system configuration
// parameter when calling OpenDatabase().
type kuzuOptions struct {
	BufferPoolSize    sobek.Value
	MaxNumThreads     sobek.Value
	EnableCompression sobek.Value
	ReadOnly          sobek.Value
	MaxDbSize         sobek.Value
}

// Handles options application for overriding the default system
// configuration for kuzudb
// DefaultSystemConfig returns the default system configuration.
// The default system configuration is as follows:
// >>> BufferPoolSize: 80% of the total system memory.
// >>> MaxNumThreads: Number of CPU cores.
// >>> EnableCompression: true.
// >>> ReadOnly: false.
// >>> MaxDbSize: 0 (unlimited).
func (opts *kuzuOptions) apply() kuzu.SystemConfig {
	kuzuConfig := kuzu.DefaultSystemConfig()

	if opts.BufferPoolSize != nil {
		kuzuConfig.BufferPoolSize = uint64(opts.BufferPoolSize.ToInteger())
	}

	if opts.MaxNumThreads != nil {
		kuzuConfig.MaxNumThreads = uint64(opts.MaxNumThreads.ToInteger())
	}

	if opts.EnableCompression != nil {
		kuzuConfig.EnableCompression = bool(opts.EnableCompression.ToBoolean())
	}

	if opts.ReadOnly != nil {
		kuzuConfig.ReadOnly = bool(opts.ReadOnly.ToBoolean())
	}

	if opts.MaxDbSize != nil {
		kuzuConfig.MaxDbSize = uint64(opts.MaxDbSize.ToInteger())
	}

	return kuzuConfig
}

type DatabaseConnection struct {
	conn *kuzu.Connection
	ctx  func() context.Context
}

func (kmod *KuzuModule) OpenConnection(databasePath string, opts *kuzuOptions) (*DatabaseConnection, error) {
	kuzuConfig := opts.apply()
	db, err := kuzu.OpenDatabase(databasePath, kuzuConfig)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	conn, err := kuzu.OpenConnection(db)
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{conn: conn, ctx: kmod.vu.Context}, nil
}

func (kuzuConn *DatabaseConnection) Query(queryString string) (*kuzu.QueryResult, error) {
	queryResult, err := kuzuConn.conn.Query(queryString)
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}
