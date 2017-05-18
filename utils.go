package nest

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func isIndex(f string) bool {
	i, err := strconv.Atoi(f)
	if err != nil || i < 0 {
		return false
	}
	return true
}

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

func pathComplete(fields []string) bool {
	return fields[0] == "$"
}

// retain structure -- *
func mergeForEach(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	for i := 0; i < v.Len(); i++ {
		if err := get(nextField(fields), v.Index(i), dst, keys...); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
	}
	return nil
}

// open structure -- .
func getForEach(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	for i := 0; i < v.Len(); i++ {
		d := reflect.New(dst.Type().Elem()).Elem()
		if err := get(nextField(fields), v.Index(i), d, keys...); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
		dst.Set(reflect.Append(dst, d))
	}
	return nil
}

func mergeForEachMap(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	for _, val := range v.MapKeys() {
		if err := get(nextField(fields), v.MapIndex(val), dst, keys...); err != nil {
			return fmt.Errorf("error setting slice index %v - %v", val, err)
		}
	}
	return nil
}

func getForEachMap(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	for _, val := range v.MapKeys() {
		d := reflect.New(dst.Type().Elem()).Elem()
		if err := get(nextField(fields), v.MapIndex(val), d, keys...); err != nil {
			return fmt.Errorf("error setting map key %s - %v", val, err)
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
