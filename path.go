package jsonast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Path []PathElement

func NewPath(elems ...PathElement) Path {
	return Path(elems)
}

func (p Path) AsPath() string {
	var ss []string
	// index -> index, e.g. [1][2]
	// index -> simple, e.g. [1].some
	// simple -> index, e.g. some[1]
	// simple -> simple, e.g. some.thing
	for _, x := range p {
		s := x.AsPath()
		switch x.(type) {
		case SimplePath:
			ss = append(ss, ".")
			ss = append(ss, s)
		default:
			ss = append(ss, s)
		}
	}

	r := strings.Join(ss, "")
	if strings.HasPrefix(r, ".") {
		return r
	}
	return "." + r
}

func (p Path) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.AsPath())
}

func (p Path) Add(elem PathElement) Path {
	return append(p, elem)
}

type PathElement interface {
	AsPath() string
}

type SimplePath string

func NewSimplePath[T StringLike](s T) SimplePath {
	return SimplePath(s)
}

func (p SimplePath) AsPath() string { return string(p) }

type IndexPath struct {
	IndexInt    int
	IndexString string
}

func NewIndexPath[I StringOrIntLike](index I) *IndexPath {
	x := &IndexPath{}

	iv := reflect.ValueOf(index)
	switch iv.Kind() {
	case reflect.Int:
		x.IndexInt = int(iv.Int())
	default:
		x.IndexString = iv.String()
	}
	return x
}

func (p IndexPath) AsPath() string {
	if p.IndexString != "" {
		return fmt.Sprintf(`["%s"]`, p.IndexString)
	}
	return fmt.Sprintf(`[%d]`, p.IndexInt)
}
