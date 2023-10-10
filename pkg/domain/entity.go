package domain

import "sync"

type Config struct {
	Port string `mapstructure:"PORT"`
}

type KeyValue struct {
	Value     interface{}
	Expiry    string
	Condition string
}

type Que struct {
	Data map[string][]string
	M    sync.Mutex
}
