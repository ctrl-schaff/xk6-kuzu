package kuzu

import (
	"go.k6.io/k6/js/modules"
)

func (*extModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &module{vu}
}

type extModule struct{}

// module represents an instance of the JavaScript module for every VU.
type module struct {
	vu modules.VU
}

// Exports is representation of ESM exports of a module.
func (m *module) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]any{
			"kuzu": m.kuzu,
		},
	}
}

var _ modules.Module = (*extModule)(nil)
