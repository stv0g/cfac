package cfac

import (
	"fmt"
)

var (
	measurables = map[string]NewMeasurable{}
)

func RegisterMeasurable(name string, ctor NewMeasurable) {
	measurables[name] = ctor
}

func GetMeasurable(name string) (NewMeasurable, error) {
	if f, ok := measurables[name]; ok {
		return f, nil
	} else {
		return nil, fmt.Errorf("unknown source: %s", name)
	}
}
