package kuzu

import (
	"fmt"

	"github.com/ctrl-schaff/xk6-kuzu/kuzu"
	"go.k6.io/k6/js/modules"
)

func init() {
	fmt.Printf("xk6-kuzu import pathing: %s\n", kuzu.ImportPath)
	newModule := kuzu.New()
	fmt.Printf("Module %v", newModule)
	modules.Register(kuzu.ImportPath, newModule)
}
