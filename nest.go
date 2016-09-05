package nest

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Get(path string, src, dst interface{}, keys ...interface{}) error {
	v := reflect.ValueOf(src)
	t := reflect.TypeOf(src)
	fmt.Println("Getting path:", path, v, t)

	d := reflect.ValueOf(dst).Elem()
	fields := strings.Split(path, "/")[1:]
	fmt.Println("Fields:", fields)

	if err := get(fields, v, d, keys...); err != nil {
		return err
	}
	return nil
}

func get(fields []string, val, dst reflect.Value, keys ...interface{}) error {
	if fields[0] == "$" {
		if dst.Type() != val.Type() {
			return fmt.Errorf("incompatible types - %s %s", val.Type(), dst.Type())
		}
		dst.Set(val)
		return nil
	}

	switch val.Kind() {
	case reflect.Slice:
		if err := processSlice(fields, val, dst, keys...); err != nil {
			return err
		}
		fmt.Println("Processed value:", dst)
	case reflect.Struct:
		if err := processStruct(fields, val, dst, keys...); err != nil {
			return err
		}
	case reflect.String, reflect.Int:
		fmt.Println("Got string", val, fields)
		dst.Set(val)
	case reflect.Ptr:
		if err := get(fields, val.Elem(), dst, keys...); err != nil {
			return err
		}
	case reflect.Map:
		if err := processMap(fields, val, dst, keys...); err != nil {
			return err
		}
	default:
		fmt.Println("Got something else", val, fields, val.Kind())
	}
	return nil
}

func processSlice(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	f := fields[0]
	switch true {
	case f == "*" || f == ".":
		switch f {
		case "*":
			if dst.Kind() != reflect.Slice {
				return fmt.Errorf("incompatible types:  %v, %v", v.Type(), dst.Type())
			}
			return getForEach(fields, v, dst, keys...)
		case ".":
			return getElemForEach(fields, v, dst, keys...)
		}
		return nil
	case isIndex(f):
		s, _ := strconv.Atoi(f)
		if s >= v.Len() {
			return fmt.Errorf("slice index exeeds length: %d", v.Len())
		}
		if err := get(nextField(fields), v.Index(s), dst, keys...); err != nil {
			return fmt.Errorf("error getting slice index %d - %v", s, err)
		}
		return nil
	}
	return fmt.Errorf("incorrect slice indexing: %v", fields[0])
}

func processStruct(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	val := v.FieldByName(fields[0])
	if !val.IsValid() {
		return fmt.Errorf("invalid field name: %v", fields[0])
	}

	if err := get(nextField(fields), val, dst, keys...); err != nil {
		return fmt.Errorf("error getting struct field %v - %v", fields[0], err)
	}
	return nil
}

func processMap(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	f := fields[0]
	var k reflect.Value
	if s, err := strconv.Atoi(f); err == nil {
		fmt.Printf("using searate key: %v\n", keys)
		if s >= len(keys) {
			return fmt.Errorf("not enough keys: %d -- %v", s, keys)
		}
		k = reflect.ValueOf(keys[s])
		if s == len(keys)-1 {
			keys = keys[:s]
		} else {
			keys = append(keys[:s], keys[s+1:])
		}
	} else {
		k = reflect.ValueOf(f)
	}
	if v.Type().Key() != k.Type() {
		return fmt.Errorf("mismatched map keys %v - %v", v.Type().Key(), k.Type())
	}
	val := v.MapIndex(k)
	if !val.IsValid() {
		return fmt.Errorf("keys not found %v", k.Type())
	}
	return get(nextField(fields), val, dst, keys...)
}
