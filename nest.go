/*
  nest is a library that provides an easy way to retreive fields from deeply nested
  data structures. Structure here can be of any Go type that can store data (slice,
  maps, structs, primitives).

  Fields are accessed by using a specific path string. The path string contains
  accessors for the nested structures seperated by '/'.

  Accessors can be of two types:
    1. Direct field/index accessor.
    2. Special types.

  Each path string must begin with a '/'. This / is the root structure. Simply getting
  the path '/' returns the original structure. e.g.:

  var dst = []int{}
  var src = []int{1, 2, 3}

  Get("/", src, &dst) // dst == []int{1, 2, 3} == src -- deep copy

  1. Direct field/index accessors:

  After the root path, direct accessor can be used. It can be an index, in case of a slice,
  or a field name/map key. e.g.:

  var dst int
  Get("/1", src, &dst) // dst == 2 == src[1]

  type Course struct {
    Name      string
    Teacher   string
  }

  var src = struct{
    Name    string
    Age     int
    Marks   []int
    Courses []Course
  }{
    Name:  "Abh",
    Age:   10,
    Marks: []int{20, 19, 15}
    Courses: []Course{ {"Physics", "Prof. P"} {"Chemistry", "Prof. C"} }
  }

  var dst string
  Get("/Name", src, &dst) // dst == "Abh" == src.Name

  var dst int
  Get("/Marks/2", src, &dst) // dst == 15 == src.Marks[2]

  var dst string
  Get("/Courses/0/Teacher", src, &dst) // dst == "Prof. P"


  2. Special Type Accessors: Use of only direct accessor result in single values,
     as illustrated above. To get multiple values, special accessors are required.
     There are two types of these accessors:

     1. The Dot (.) Accessor:

     The Dot accessor fetches each element in an iterable type (slice, map). For example,
     consider the following value:

     type InnerSimple struct {
       I int
       S string
     }
     v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}

     Now, to get all the values in v, we can write:

     var result []InnerSimple
     Get("/.", v, &result) // result == v

     But this is not a very useful result. But we can see that '.' fetched each value in v
     and put it in a new slice. Now consider the case when you want all the string values
     from the slice:

     var result []int
     Get("/./S", v, &result) // result == []string{"one", "two", "three"}

     Now this is a more useful result. '.' can be used multiple times in a path as long as
     it follows an iterable type in the path. In above path, '/' is a slice itself and is
     iterable. The '.' accesses each element in that slice ('/' here) and fetches the elements
     'S' field. Note that each element of the slice must either be a struct with a field named
     'S' or a map with a key 'S'.

     As another example, consider:

     type Simple struct {
       SimpleI int
       SimpleS string
       InnerSimple
       SimpleSlc []InnerSimple
     }

     v := []Simple{
       {1, "one", InnerSimple{10, "ten"}, []InnerSimple{InnerSimple{1, "a"}, InnerSimple{2, "b"}}},
       {2, "two", InnerSimple{20, "twenty"}, []InnerSimple{InnerSimple{3, "c"}, InnerSimple{4, "d"}}},
       {3, "three", InnerSimple{30, "thirty"}, []InnerSimple{InnerSimple{5, "e"}, InnerSimple{6, "f"}}},
     }

     var result [][]string
     Get("/./SimpleSlc/./S", v, &result) // result ==  [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}

     In above example, the path starts at root, which is a slice ([]Simple). It then accesses
     each element of this slice. For each element, it accesses the SimpleSlc field. This
     field is also a slice and so another '.' can be used to access each element of SimpleSlc.
     After the second '.', 'S' field of each element of SimpleSlc is accessed.

     Note the type of the result. It is of [][]string type. That's because each '.' fills up
     exactly one slice. Inner '.' works on a single SimpleSlc and returns a slice containing
     all the 'S' values. The outer '.' collect these values in another slice. Hence [][]string.

     But what if we wanted to merge the inner slices together. That's where our second
     accessor type comes into play.

     2. The Star (*) Accessor:

     There are two important rules associated with this accessor.

     a) Star accessor can only be used to replace a '.' in the path.
     b) Star accessor cannot be the last accessor (or only) in a path. I.e., path '/A/B/* /C'
        is illegal.

     What '*' does is that it breaks the structure or returned result. If in above path
     "/./SimpleSlc/./S", we replace the first '.' with a star, we get the following result:

     var result [][]string
     Get("/* /SimpleSlc/./S", v, &result) // result == []string{"a", "b", "c", "d", "e", "f"}

     As evident, '*' merges together the outer slice. This is why '*' must always be used
     before a '.', because it must have some result to merge together.

     Many more examples can be found in the test files, which also include map accesses.


  TODO: Add data updates through path.
        Add custom function to process accessed data.
        Add interface similar to Marshalling/Unmarshalling similar to encoding.

*/
package nest

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Get gets the values specified by the path from the source value. The retreived values are
// stored in the destination value. The destination value must be settable.
// TODO: Utilize keys.
func Get(path string, src, dst interface{}, keys ...interface{}) error {
	v := reflect.ValueOf(src)

	d := reflect.ValueOf(dst).Elem()
	fields := strings.Split(path, "/")[1:]

	if err := get(fields, v, d, keys...); err != nil {
		return err
	}
	return nil
}

func get(fields []string, val, dst reflect.Value, keys ...interface{}) error {
	if pathComplete(fields) {
		if dst.Type() != val.Type() {
			return fmt.Errorf("get: incompatible types - src: %s | dst: %s", val.Type(), dst.Type())
		}
		dst.Set(val)
		return nil
	}

	var err error
	switch val.Kind() {
	case reflect.String, reflect.Int:
		dst.Set(val)
	case reflect.Slice:
		err = processSlice(fields, val, dst, keys...)
	case reflect.Struct:
		err = processStruct(fields, val, dst, keys...)
	case reflect.Map:
		err = processMap(fields, val, dst, keys...)
	case reflect.Ptr, reflect.Interface:
		err = get(fields, val.Elem(), dst, keys...)
	default:
		// TODO: Handle default better.
		err = fmt.Errorf("unknown value type, cannot process: %s", val.Kind())
	}
	return err
}

func processSlice(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	f := fields[0]
	switch true {
	case f == ".":
		return getForEach(fields, v, dst, keys...)
		break
	case f == "*":
		if dst.Kind() != reflect.Slice {
			return fmt.Errorf("processSlice: incompatible types: src: %v | dst: %v", v.Type(), dst.Type())
		}
		return mergeForEach(fields, v, dst, keys...)
		break
	case isIndex(f):
		s, _ := strconv.Atoi(f)
		if s >= v.Len() {
			return fmt.Errorf("processSlice: slice index exeeds length: %d", v.Len())
		}
		if err := get(nextField(fields), v.Index(s), dst, keys...); err != nil {
			return fmt.Errorf("processSlice: error getting slice index %d - %v", s, err)
		}
		break
	default:
		return fmt.Errorf("processSlice: incorrect slice indexing: %v", fields[0])
	}
	return nil
}

func processStruct(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	val := v.FieldByName(fields[0])
	if !val.IsValid() {
		return fmt.Errorf("processStruct: invalid field name: %v", fields[0])
	}

	if err := get(nextField(fields), val, dst, keys...); err != nil {
		return fmt.Errorf("processStruct: error getting struct field %v - %v", fields[0], err)
	}
	return nil
}

func processMap(fields []string, v, dst reflect.Value, keys ...interface{}) error {
	f := fields[0]
	var k reflect.Value
	switch true {
	case isIndex(f):
		s, _ := strconv.Atoi(f)
		if s >= len(keys) {
			return fmt.Errorf("not enough keys: %d -- %v", s, keys)
		}
		k = reflect.ValueOf(keys[s])
		if s == len(keys)-1 {
			keys = keys[:s]
		} else {
			keys = append(keys[:s], keys[s+1:])
		}
	case f == ".":
		return getForEachMap(fields, v, dst, keys...)
	case f == "*":
		if dst.Kind() != reflect.Slice {
			return fmt.Errorf("processMap: incompatible types: src: %v | dst: %v", v.Type(), dst.Type())
		}
		return mergeForEachMap(fields, v, dst, keys...)
	default:
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
