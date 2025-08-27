package kuzu

import (
	"go.k6.io/k6/js/modules"
)

type rootModule struct{}

func (*rootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &module{vu}
}

type module struct {
	vu modules.VU
}

func (m *module) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]any{
			"kuzu": m,
		},
	}
}

var _ modules.Module = (*rootModule)(nil)
