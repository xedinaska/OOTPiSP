package plugins

import (
	"fmt"
	"plugin"
)

func Load(name string) (*plugin.Plugin, error) {
	return plugin.Open(fmt.Sprintf("plugins/bin/%s.so", name))
}
