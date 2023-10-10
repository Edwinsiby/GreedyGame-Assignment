package delivery

type keyValueInput struct {
	Key   string
	Value KeyValue
}

type KeyValue struct {
	Value     interface{}
	Expiry    int
	Condition string
}

type key struct {
	Key string
}

type QueInput struct {
	Key   string
	Value []string
}

type BQPopInput struct {
	Key  string
	Time int
}
