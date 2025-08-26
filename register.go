package kuzu

import "go.k6.io/k6/js/modules"

const importPath = "k6/x/kuzu"

func init() {
	modules.Register(importPath, new(extModule))
}
