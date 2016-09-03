package nest

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Get(path string, src, dst interface{}) error {
	// if fields, err := parsePath(path); err != nil {
	// 	return fmt.Errorf("illegal path modifiers: %v", err)
	// }

	v := reflect.ValueOf(src)
	t := reflect.TypeOf(src)
	fmt.Println("Getting path:", path, v, t)

	d := reflect.ValueOf(dst).Elem()
	fields := strings.Split(path, "/")[1:]
	fmt.Println("Fields:", fields)

	if err := get(fields, v, d); err != nil {
		return err
	}
	return nil
}

func get(fields []string, val, dst reflect.Value) error {
	if fields[0] == "$" {
		if dst.Type() != val.Type() {
			return fmt.Errorf("incompatible types - %s %s", dst.Type(), val.Type())
		}
		dst.Set(val)
		return nil
	}

	switch val.Kind() {
	case reflect.Slice:
		return processSlice(fields, val, dst)
		// fmt.Println("Processed value:", dst)
	case reflect.Struct:
		return processStruct(fields, val, dst)
	case reflect.Map:
		return processMap(fields, val, dst)
	case reflect.String, reflect.Int:
		fmt.Println("Got string", val, fields)
		dst.Set(val)
	case reflect.Ptr:
		return get(fields, val.Elem(), dst)
	default:
		fmt.Println("Got something else", val, fields, val.Kind())
	}
	return nil
}

func processSlice(fields []string, v, dst reflect.Value) error {
	f, mod := getField(fields[0])
	s, err := strconv.Atoi(f)
	switch true {
	case f == "*":
		if dst.Kind() != reflect.Slice {
			return fmt.Errorf("incompatible types:  %v, %v", v.Type(), dst.Type())
		}
		switch mod {
		case "-":
			return getForEach(fields, v, dst)
		default:
			return getElemForEach(fields, v, dst)
		}
		return nil
	case err == nil:
		if s >= v.Len() {
			return fmt.Errorf("slice index exeeds length: %d", v.Len())
		}
		if err := get(nextField(fields), v.Index(s), dst); err != nil {
			return fmt.Errorf("error getting slice index %d - %v", s, err)
		}
		return nil
	}
	return fmt.Errorf("incorrect slice indexing: %v", fields[0])
}

func processStruct(fields []string, v, dst reflect.Value) error {
	val := v.FieldByName(fields[0])
	if !val.IsValid() {
		return fmt.Errorf("invalid field name: %v", fields[0])
	}

	if err := get(nextField(fields), val, dst); err != nil {
		return fmt.Errorf("error getting struct field %v - %v", fields[0], err)
	}
	return nil
}

func processMap(fields []string, v, dst reflect.Value) error {
	f, _ := getField(fields[0])
	k := reflect.ValueOf(f)
	val := v.MapIndex(k)
	return get(nextField(fields), val, dst)
}
