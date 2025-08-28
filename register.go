package kuzu

import (
	"github.com/ctrl-schaff/xk6-kuzu/kuzu"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/kuzu", kuzu.NewModule())
}
