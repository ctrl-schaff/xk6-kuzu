package kuzu

import (
	"github.com/grafana/sobek"
	"github.com/kuzudb/go-kuzu"
)

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
