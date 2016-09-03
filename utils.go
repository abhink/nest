package nest

import (
	"fmt"
	"reflect"
)

func getForEach(fields []string, v, dst reflect.Value) error {
	for i := 0; i < v.Len(); i++ {
		if err := get(nextField(fields), v.Index(i), dst); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
	}
	return nil
}

func getElemForEach(fields []string, v, dst reflect.Value) error {
	for i := 0; i < v.Len(); i++ {
		d := reflect.New(dst.Type().Elem()).Elem()
		if err := get(nextField(fields), v.Index(i), d); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
		dst.Set(reflect.Append(dst, d))
	}
	return nil
}

func set(val, dst reflect.Value) error {
	if dst.Type() != val.Type() {
		return fmt.Errorf("incompatible types - %s %s", dst.Type(), val.Type())
	}
	return nil
}
