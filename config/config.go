// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	Table  []string      `config:"table"`
}

var DefaultConfig = Config{
	Period: 10 * time.Second,
	Table:	[]string{"system.local"},
}
