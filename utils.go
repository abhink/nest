package nest

import (
	"fmt"
	"reflect"
	"strings"
)

func getField(f string) (string, string) {
	fs := strings.Split(f, ":")
	if len(fs) == 1 {
		return fs[0], ""
	}
	return fs[0], fs[1]
}

func nextField(fields []string) []string {
	if len(fields) == 1 {
		return []string{"$"}
	}
	return fields[1:]
}

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
