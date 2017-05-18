// Tests/Example for reference. In each test function, 'v' is the src value
// while 'r' is the expected result.
package nest

import (
	"reflect"
	"sort"
	"testing"
)

type InnerSimple struct {
	I int
	S string
}

type Simple struct {
	SimpleI int
	SimpleS string
	InnerSimple
	SimpleSlc []InnerSimple
}

func TestSlice(t *testing.T) {

	v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
	r := []string{"one", "two", "three"}

	var result []string

	if err := Get("/./S", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}

	intResult := []int{}

	if err := Get("/./I", v, &intResult); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(intResult, []int{1, 2, 3}) {
		t.Errorf("unexpected ouput: %v -- %T", intResult, intResult)
	}

	intrm := []int{}
	for _, o := range v {
		intrm = append(intrm, o.I)
	}
	if !reflect.DeepEqual(intrm, []int{1, 2, 3}) {
		t.Errorf("unexpected ouput: %v -- %T", intrm, intrm)
	}
}

func TestSliceNested(t *testing.T) {
	v := []Simple{
		{1, "one", InnerSimple{10, "ten"}, nil},
		{2, "two", InnerSimple{20, "twenty"}, nil},
		{3, "three", InnerSimple{30, "thirty"}, nil},
	}
	r := []InnerSimple{{10, "ten"}, {20, "twenty"}, {30, "thirty"}}

	var result []InnerSimple

	if err := Get("/./InnerSimple", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}
}

func TestSliceNestedField(t *testing.T) {
	v := []Simple{
		{1, "one", InnerSimple{10, "ten"}, nil},
		{2, "two", InnerSimple{20, "twenty"}, nil},
		{3, "three", InnerSimple{30, "thirty"}, nil},
	}
	r := []int{10, 20, 30}

	var result []int

	if err := Get("/./InnerSimple/I", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}
}

func TestSliceNestedSlice(t *testing.T) {
	v := []Simple{
		{1, "one", InnerSimple{10, "ten"}, []InnerSimple{InnerSimple{1, "a"}, InnerSimple{2, "b"}}},
		{2, "two", InnerSimple{20, "twenty"}, []InnerSimple{InnerSimple{3, "c"}, InnerSimple{4, "d"}}},
		{3, "three", InnerSimple{30, "thirty"}, []InnerSimple{InnerSimple{5, "e"}, InnerSimple{6, "f"}}},
	}
	r := [][]InnerSimple{
		[]InnerSimple{InnerSimple{1, "a"}, InnerSimple{2, "b"}},
		[]InnerSimple{InnerSimple{3, "c"}, InnerSimple{4, "d"}},
		[]InnerSimple{InnerSimple{5, "e"}, InnerSimple{6, "f"}},
	}

	var result [][]InnerSimple

	if err := Get("/./SimpleSlc", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}
}

func TestSliceNestedSliceDot(t *testing.T) {
	v := []Simple{
		{1, "one", InnerSimple{10, "ten"}, []InnerSimple{InnerSimple{1, "a"}, InnerSimple{2, "b"}}},
		{2, "two", InnerSimple{20, "twenty"}, []InnerSimple{InnerSimple{3, "c"}, InnerSimple{4, "d"}}},
		{3, "three", InnerSimple{30, "thirty"}, []InnerSimple{InnerSimple{5, "e"}, InnerSimple{6, "f"}}},
	}
	r := [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}

	var result [][]string

	if err := Get("/./SimpleSlc/./S", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}
}

func TestSliceNestedSliceStar(t *testing.T) {
	v := []Simple{
		{1, "one", InnerSimple{10, "ten"}, []InnerSimple{InnerSimple{1, "a"}, InnerSimple{2, "b"}}},
		{2, "two", InnerSimple{20, "twenty"}, []InnerSimple{InnerSimple{3, "c"}, InnerSimple{4, "d"}}},
		{3, "three", InnerSimple{30, "thirty"}, []InnerSimple{InnerSimple{5, "e"}, InnerSimple{6, "f"}}},
	}
	r := []string{"a", "b", "c", "d", "e", "f"}

	var result []string

	if err := Get("/*/SimpleSlc/./S", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}
}

func TestSliceFlatten(t *testing.T) {
	type InnerSimpleSlc struct {
		InnerSimple
		Slc []int
	}
	v := []InnerSimpleSlc{
		{InnerSimple{1, "one"}, []int{1, 2}},
		{InnerSimple{2, "two"}, []int{2, 4}},
		{InnerSimple{3, "three"}, []int{3, 6}},
	}
	r := 4

	var result int

	if err := Get("/1/Slc/1", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}
}

func TestGetSliceIndex(t *testing.T) {
	v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
	r := 2

	var result int

	if err := Get("/1/I", v, &result); err != nil {
		t.Errorf("error Getting index: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestGetSliceIndexAtom(t *testing.T) {
	v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
	r := InnerSimple{2, "two"}

	var result InnerSimple

	if err := Get("/1", v, &result); err != nil {
		t.Errorf("error Getting index: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestGetSliceIndexMismatchType(t *testing.T) {
	v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
	var result int

	if err := Get("/1", v, &result); err == nil {
		t.Errorf("should fail on mismatched types")
	}
}

func TestGetSliceInvalidIndex(t *testing.T) {
	v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
	var result int

	if err := Get("/InvalidIndex", v, &result); err == nil {
		t.Errorf("should fail on invalid index")
	}
}

func TestGetStructField(t *testing.T) {
	v := InnerSimple{1, "one"}
	r := "one"
	var result string

	if err := Get("/S", v, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result != r {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplextAtom(t *testing.T) {
	r := "AB1D1E1"
	var result string

	if err := Get("/Bslc/0/Dslc/0/Eslc/0/Estr", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result != r {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSlice(t *testing.T) {
	r := [][]string{
		[]string{"AB1D1", "AB1D2", "AB1D3"},
		[]string{"AB2D1", "AB2D2", "AB2D3"},
	}
	var result = [][]string{}

	if err := Get("/Bslc/./Dslc/./Dstr", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceFlatten(t *testing.T) {
	r := []string{"AB1D1", "AB1D2", "AB1D3", "AB2D1", "AB2D2", "AB2D3"}
	var result = []string{}

	if err := Get("/Bslc/*/Dslc/./Dstr", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceInt(t *testing.T) {
	r := []int{11, 12, 13}
	var result = []int{}

	if err := Get("/Bslc/0/Dslc/0/Eslc/./Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceIntFlatten(t *testing.T) {
	r := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	var result = []int{}

	if err := Get("/Bslc/*/Dslc/*/Eslc/./Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceIntFlattenPartial1(t *testing.T) {
	r := [][]int{
		[]int{11, 12, 13, 14, 15, 16, 17, 18, 19},
		[]int{21, 22, 23, 24, 25, 26, 27, 28, 29},
	}
	var result = [][]int{}

	if err := Get("/Bslc/./Dslc/*/Eslc/./Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceIntFlattenPartial2(t *testing.T) {
	r := [][]int{
		[]int{11, 12, 13}, []int{14, 15, 16}, []int{17, 18, 19},
		[]int{21, 22, 23}, []int{24, 25, 26}, []int{27, 28, 29},
	}
	var result = [][]int{}

	if err := Get("/Bslc/*/Dslc/./Eslc/./Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexValue(t *testing.T) {
	r := "AB2D2E2"
	var result = [][][]*E{}

	if err := Get("/Bslc/./Dslc/./Eslc", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result[1][1][1].Estr, r) {
		t.Errorf("unexpected ouput: %#v -- %v", result[1][1][1].Estr, r)
	}
}

func TestComplexValueFlattenB(t *testing.T) {
	r := 18
	var result = [][]*E{}

	if err := Get("/Bslc/*/Dslc/./Eslc", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result[2][1].Eint != r {
		t.Errorf("unexpected ouput: %#v -- %v", result[2][1].Eint, r)
	}

	// Below is the code that would be required to replace Get call above.
	// for _, b := range Av {
	// 	for _, d := range b {
	// 		for _, e := range d {
	// 			es = append(es, e.Estr...)
	// 		}
	// 	}
	// }
}

func TestComplexValueFlattenBD(t *testing.T) {
	r := 18
	var result = []*E{}

	if err := Get("/Bslc/*/Dslc/*/Eslc/.", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result[7].Eint != r {
		t.Errorf("unexpected ouput: %#v -- %v", result[8].Eint, r)
	}

	// TOFO: Add more modifiers.
	// Example:
	// Get("/Bslc/:ALL:FLATTEN/Dslc/:ALL:FLATTEN/Eslc/:ALL", Av, &result)
	// Get("/Bslc:ALL:FLATTEN/Dslc:ALL:FLATTEN/Eslc:ALL", Av, &result)
	// Get("/Bslc/:A:F/Dslc/:A:F/Eslc/:A", Av, &result)
	// Get("/Bslc:A:F/Dslc:A:F/Eslc:A", Av, &result)
}

func TestComplexValueIndex(t *testing.T) {
	r := 18
	var result = &E{}

	if err := Get("/Bslc/0/Dslc/2/Eslc/1", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result.Eint != r {
		t.Errorf("unexpected ouput: %#v -- %v", result.Eint, r)
	}
}

func TestComplexValueIndexC(t *testing.T) {
	r := "AB1CE1"
	var result string

	if err := Get("/Bslc/0/Cptr/Eptr/Estr", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result != r {
		t.Errorf("unexpected ouput: %#v -- %v", result, r)
	}
}

func TestComplexValueSliceC(t *testing.T) {
	r := []int{1, 1}
	var result = []int{}

	if err := Get("/Bslc/./Cptr/Eptr/Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %v", result, r)
	}
}

func TestMapSimple(t *testing.T) {
	r := "222 33"
	var result string

	if err := Get("/Bslc/0/Dslc/1/Eslc/2/Emap/k3", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %v", result, r)
	}
}

func TestMapSlice(t *testing.T) {
	r := []string{
		"111 13", "111 23", "111 33",
		"222 13", "222 23", "222 33",
		"333 13", "333 23", "333 33",
	}
	var result = []string{}

	if err := Get("/Bslc/0/Dslc/*/Eslc/./Emap/k3", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapSliceIndex(t *testing.T) {
	r := []string{
		"111 13", "111 23", "111 33",
		"222 13", "222 23", "222 33",
		"333 13", "333 23", "333 33",
	}
	var result = []string{}

	if err := Get("/Bslc/0/Dslc/*/Eslc/./Emap/0", Av, &result, "k3"); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapSliceIndexComplex(t *testing.T) {
	r := []*E{e2, e5}
	var result = []*E{}

	if err := Get("/Bslc/./Bmap/0", Av, &result, mk2); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestSliceSlice(t *testing.T) {
	r := []string{
		"aaa", "sss", "ddd", "zzz", "xxx", "ccc", "qqq", "www", "eee", "rrr", "fff",
	}

	var result = []string{}

	if err := Get("/S/*/.", Bv, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestSliceSliceNested(t *testing.T) {
	r := [][]string{
		[]string{"aaa", "sss", "ddd"},
		[]string{"zzz", "xxx", "ccc"},
		[]string{"qqq", "www", "eee", "rrr", "fff"},
	}

	var result = [][]string{}

	if err := Get("/S/.", Bv, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestSliceOnly(t *testing.T) {
	v := [][]string{
		[]string{"aaa", "sss", "ddd"},
		[]string{"zzz", "xxx", "ccc"},
		[]string{"qqq", "www", "eee", "rrr", "fff"},
	}
	r := []string{
		"zzz", "xxx", "ccc",
	}

	var result = []string{}

	if err := Get("/1/.", v, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapSimpleDot(t *testing.T) {
	m := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
	r := []int{1, 2, 3, 4}
	var result []int

	if err := Get("/.", m, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	sort.Ints(result)
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapSimpleDotField(t *testing.T) {
	m := map[string]InnerSimple{"A": {1, "one"}, "B": {2, "two"}, "C": {3, "three"}}
	r := []int{1, 2, 3}
	var result []int

	if err := Get("/./I", m, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	sort.Ints(result)
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapSimpleDotSlice(t *testing.T) {
	m := map[string][]InnerSimple{
		"A": []InnerSimple{{11, "oneA"}, {12, "twoA"}, {13, "threeA"}},
		"B": []InnerSimple{{21, "oneB"}, {22, "twoB"}, {23, "threeB"}},
		"C": []InnerSimple{{31, "oneC"}, {32, "twoC"}, {33, "threeC"}},
	}
	r := []int{11, 12, 13, 21, 22, 23, 31, 32, 33}
	var result []int

	if err := Get("/*/./I", m, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	sort.Ints(result)
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapJSON(t *testing.T) {
	r := []interface{}{"T1A TV", "T3A", "T3B"}
	var result = make([]interface{}, 0)

	if err := Get("/category/television/*/subCategory/./subCat", JSONMap, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapJSON1(t *testing.T) {
	r := 24
	var result interface{}

	if err := Get("/category/television_warrantyPeriod", JSONMap, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapJSON2(t *testing.T) {
	r := 5
	var result interface{}

	if err := Get("/category/television/1/subCategory/1/warrantyPeriod", JSONMap, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestMapJSON3(t *testing.T) {
	r := 5
	var result interface{}

	if err := Get("/category/television/1/0/1/warrantyPeriod", JSONMap, &result, "subCategory"); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func xTestMapJSON4(t *testing.T) {
	r := 5
	var result interface{}

	if err := Get("/category/0/1/1/1/warrantyPeriod", JSONMap, &result, "subCategory"); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}
