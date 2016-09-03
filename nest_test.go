package nest

import (
	"reflect"
	"testing"
)

type InnerSimple struct {
	I int
	S string
}

func TestSlice(t *testing.T) {

	v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
	r := []string{"one", "two", "three"}

	var result []string

	if err := Get("/*/S", v, &result); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %T", result, result)
	}

	intr := []int{}

	if err := Get("/*/I", v, &intr); err != nil {
		t.Errorf("error Getting value: %v", err)
	}

	if !reflect.DeepEqual(intr, []int{1, 2, 3}) {
		t.Errorf("unexpected ouput: %v -- %T", intr, intr)
	}

	intrm := []int{}
	for _, o := range v {
		intrm = append(intrm, o.I)
	}
	if !reflect.DeepEqual(intrm, []int{1, 2, 3}) {
		t.Errorf("unexpected ouput: %v -- %T", intrm, intrm)
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

	if err := Get("/Bslc/*/Dslc/*/Dstr", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceFlatten(t *testing.T) {
	r := []string{"AB1D1", "AB1D2", "AB1D3", "AB2D1", "AB2D2", "AB2D3"}
	var result = []string{}

	if err := Get("/Bslc/*:-/Dslc/*/Dstr", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceInt(t *testing.T) {
	r := []int{11, 12, 13}
	var result = []int{}

	if err := Get("/Bslc/0/Dslc/0/Eslc/*/Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexSliceIntFlatten(t *testing.T) {
	r := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	var result = []int{}

	if err := Get("/Bslc/*:-/Dslc/*:-/Eslc/*/Eint", Av, &result); err != nil {
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

	if err := Get("/Bslc/*/Dslc/*:-/Eslc/*/Eint", Av, &result); err != nil {
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

	if err := Get("/Bslc/*:-/Dslc/*/Eslc/*/Eint", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %v -- %v", result, r)
	}
}

func TestComplexValue(t *testing.T) {
	r := "AB2D2E2"
	var result = [][][]*E{}

	if err := Get("/Bslc/*/Dslc/*/Eslc", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result[1][1][1].Estr, r) {
		t.Errorf("unexpected ouput: %#v -- %v", result[1][1][1].Estr, r)
	}
}

func TestComplexValueFlattenB(t *testing.T) {
	r := 18
	var result = [][]*E{}

	if err := Get("/Bslc/*:-/Dslc/*/Eslc", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result[2][1].Eint != r {
		t.Errorf("unexpected ouput: %#v -- %v", result[2][1].Eint, r)
	}

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

	if err := Get("/Bslc/*:-/Dslc/*:-/Eslc/*", Av, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if result[7].Eint != r {
		t.Errorf("unexpected ouput: %#v -- %v", result[8].Eint, r)
	}
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

	if err := Get("/Bslc/*/Cptr/Eptr/Eint", Av, &result); err != nil {
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

	if err := Get("/Bslc/0/Dslc/*:-/Eslc/*/Emap/k3", Av, &result); err != nil {
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

	if err := Get("/S/*:-/*", Bv, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}

func TestSliceSliceNoFlat(t *testing.T) {
	r := [][]string{
		[]string{"aaa", "sss", "ddd"},
		[]string{"zzz", "xxx", "ccc"},
		[]string{"qqq", "www", "eee", "rrr", "fff"},
	}

	var result = [][]string{}

	if err := Get("/S/*", Bv, &result); err != nil {
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
		"aaa", "sss", "ddd", "zzz", "xxx", "ccc", "qqq", "www", "eee", "rrr", "fff",
	}

	var result = []string{}

	if err := Get("/*:-/*", v, &result); err != nil {
		t.Errorf("error Getting field: %v", err)
	}
	if !reflect.DeepEqual(result, r) {
		t.Errorf("unexpected ouput: %#v -- %#v", result, r)
	}
}
