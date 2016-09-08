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

// retain structure -- *
func getForEach(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	for i := 0; i < v.Len(); i++ {
		if err := get(nextField(fields), v.Index(i), dst, keys...); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
	}
	return nil
}

// open structure -- .
func getElemForEach(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	for i := 0; i < v.Len(); i++ {
		d := reflect.New(dst.Type().Elem()).Elem()
		if err := get(nextField(fields), v.Index(i), d, keys...); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
		dst.Set(reflect.Append(dst, d))
	}
	return nil
}

func getElemForEachMap(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	var vals []reflect.Value
	for _, val := range v.MapKeys() {
		vals = append(vals, v.MapIndex(val))
	}
	mapValues := reflect.ValueOf(vals)
	for i := 0; i < v.Len(); i++ {
		if err := get(nextField(fields), mapValues, dst, keys...); err != nil {
			return fmt.Errorf("error setting slice index %d - %v", i, err)
		}
	}
	return nil
}

func set(val, dst reflect.Value) error {
	if dst.Type() != val.Type() {
		return fmt.Errorf("incompatible types - %s %s", dst.Type(), val.Type())
	}
	return nil
}

/*

  err := From(JSONMap, &result).Get("category").Get("television").ForAll().
    Get("subCategory").ForEach().Get("subCat")
  if err != nil {
    t.Errorf("error Getting field: %v", err)
  }

*/

type partial struct {
	val reflect.Value
	dst reflect.Value
	err error
}

func From(src, dst interface{}) *partial {
	from := reflect.ValueOf(src)
	to := reflect.ValueOf(dst).Elem()

	return &partial{
		from,
		to,
		nil,
	}
}

func (p *partial) Get(key interface{}) *partial {

}
