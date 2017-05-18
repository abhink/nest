package nest

type A struct {
	Astr string
	Bslc []*B
}

type B struct {
	Bstr string
	Cptr *C
	Dslc []*D
	Bmap map[*mkey]*E
}

type C struct {
	Cstr string
	Eptr *E
	Cmap me
}

type D struct {
	Dstr string
	Eslc []*E
}

type E struct {
	Eint int
	Estr string
	Emap ms
}

type ms map[string]string

type me map[int]*E

var e1 = &E{111, "OUT1", ms{"k1": "111 1", "k2": "111 1", "k3": "111 1"}}
var e2 = &E{112, "OUT2", ms{"k1": "112 1", "k2": "112 1", "k3": "112 1"}}
var e3 = &E{113, "OUT3", ms{"k1": "113 1", "k2": "113 1", "k3": "113 1"}}
var e4 = &E{114, "OUT4", ms{"k1": "114 1", "k2": "114 1", "k3": "114 1"}}
var e5 = &E{115, "OUT5", ms{"k1": "115 1", "k2": "115 1", "k3": "115 1"}}

type mkey struct {
	I int
	S string
}

var mk1 = &mkey{1, "mkey1"}
var mk2 = &mkey{2, "mkey2"}
var mk3 = &mkey{3, "mkey3"}
var mk4 = &mkey{4, "mkey4"}
var mk5 = &mkey{5, "mkey5"}

var Av = &A{
	Astr: "A",
	Bslc: []*B{
		&B{
			"AB1",
			&C{"AB1C", &E{1, "AB1CE1", ms{"c": "1"}}, me{1: e1, 3: e3}},
			[]*D{
				&D{
					"AB1D1", []*E{
						&E{11, "AB1D1E1", ms{"ki1": "111 11", "k2": "111 12", "k3": "111 13"}},
						&E{12, "AB1D1E2", ms{"ki1": "111 21", "k2": "111 22", "k3": "111 23"}},
						&E{13, "AB1D1E3", ms{"ki1": "111 31", "k2": "111 32", "k3": "111 33"}},
					},
				}, &D{
					"AB1D2", []*E{
						&E{14, "AB1D2E1", ms{"ki1": "222 11", "k2": "222 12", "k3": "222 13"}},
						&E{15, "AB1D2E2", ms{"ki1": "222 21", "k2": "222 22", "k3": "222 23"}},
						&E{16, "AB1D2E3", ms{"ki1": "222 31", "k2": "222 32", "k3": "222 33"}},
					},
				}, &D{
					"AB1D3", []*E{
						&E{17, "AB1D3E1", ms{"ki1": "333 11", "k2": "333 12", "k3": "333 13"}},
						&E{18, "AB1D3E2", ms{"ki1": "333 21", "k2": "333 22", "k3": "333 23"}},
						&E{19, "AB1D3E3", ms{"ki1": "333 31", "k2": "333 32", "k3": "333 33"}},
					},
				},
			},
			map[*mkey]*E{mk1: e1, mk2: e2},
		}, &B{
			"AB2",
			&C{"AB2C", &E{1, "AB2CE1", ms{"c": "1"}}, me{1: e1, 4: e4, 5: e5, 2: e1}},
			[]*D{
				&D{
					"AB2D1", []*E{
						&E{21, "AB2D1E1", ms{"ki1": "444 11", "k2": "444 12", "k3": "444 13"}},
						&E{22, "AB2D1E2", ms{"ki1": "444 21", "k2": "444 22", "k3": "444 23"}},
						&E{23, "AB2D1E3", ms{"ki1": "444 31", "k2": "444 32", "k3": "444 33"}},
					},
				}, &D{
					"AB2D2", []*E{
						&E{24, "AB2D2E1", ms{"ki1": "444 11", "k2": "444 12", "k3": "444 13"}},
						&E{25, "AB2D2E2", ms{"ki1": "444 21", "k2": "444 22", "k3": "444 23"}},
						&E{26, "AB2D2E3", ms{"ki1": "444 31", "k2": "444 32", "k3": "444 33"}},
					},
				}, &D{
					"AB2D3", []*E{
						&E{27, "AB2D3E1", ms{"ki1": "444 11", "k2": "444 12", "k3": "444 13"}},
						&E{28, "AB2D3E2", ms{"ki1": "444 21", "k2": "444 22", "k3": "444 23"}},
						&E{29, "AB2D3E3", ms{"ki1": "444 31", "k2": "444 32", "k3": "444 33"}},
					},
				},
			},
			map[*mkey]*E{mk1: e4, mk2: e5, mk3: e3},
		},
	},
}

var Bv = &struct {
	S [][]string
}{
	S: [][]string{
		[]string{"aaa", "sss", "ddd"},
		[]string{"zzz", "xxx", "ccc"},
		[]string{"qqq", "www", "eee", "rrr", "fff"},
	},
}

type jsonm map[string]interface{}

var JSONMap = jsonm{
	"brandId": 123,
	"category": map[string]interface{}{
		"television": []interface{}{
			map[string]interface{}{
				"cat": "T1",
				"subCategory": []interface{}{
					map[string]interface{}{
						"subCat":         "T1A TV",
						"warrantyPeriod": 6,
					},
				},
				"warrantyPeriod": 12,
			},
			map[string]interface{}{
				"warrantyPeriod": 4,
				"cat":            "T3",
				"subCategory": []interface{}{
					map[string]interface{}{
						"subCat":         "T3A",
						"warrantyPeriod": 3,
					},
					map[string]interface{}{
						"subCat":         "T3B",
						"warrantyPeriod": 5,
					},
				},
			},
		},
		"radio": []interface{}{
			map[string]interface{}{
				"cat": "R1",
				"subCategory": []interface{}{
					map[string]interface{}{
						"subCat":         "R1A RAD",
						"warrantyPeriod": 16,
					},
				},
				"warrantyPeriod": 2,
			},
			map[string]interface{}{
				"warrantyPeriod": 4,
				"cat":            "R3",
				"subCategory": []interface{}{
					map[string]interface{}{
						"subCat":         "R3A",
						"warrantyPeriod": 3,
					},
					map[string]interface{}{
						"subCat":         "R3B",
						"warrantyPeriod": 5,
					},
				},
			},
		},
		"television_warrantyPeriod": 24,
	},
	"title": "BrandName",
	"_id":   "string",
}
