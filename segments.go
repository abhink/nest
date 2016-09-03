package nest

import (
	"fmt"
	"strings"
)

type ftype uint32

const (
	INDEX = ftype(1 << iota)
	NAME
	NIL
)

type mtype uint32

const (
	ALL = mtype(1 << iota)
	FLATTEN
	STRING

	ALLF = ALL | FLATTEN
)

type segment struct {
	f string
	t ftype
	m mtype
}

func parsePath(path string) ([]segment, error) {
	ps := strings.Split(path, "/")
	segments := []segment{}
	var s segment

	for i, p := range ps {
		f, m := "", 1 //processSegment(p)
		if f == "" {
			fmt.Println(f, m, s, i, p)
		}
	}
	return segments, nil
}

func getModifiers(f string) (string, []string) {
	fs := strings.Split(f, ":")
	if len(fs) == 1 {
		return fs[0], nil
	} else if len(fs) > 1 && fs[0] == "" {
		return fs[0], nil
	}
	return fs[0], fs[1:]
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
