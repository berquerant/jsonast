package jsonast

import (
	"encoding/json"

	"github.com/alecthomas/participle/v2/lexer"
)

type Marshaler struct{}

func (m Marshaler) Marshal(node Node) ([]byte, error) {
	return json.Marshal(m.convert(node, NewPath()))
}

func (m Marshaler) Build(node Node, initialPath Path) JNode {
	return m.convert(node, initialPath)
}

func (m Marshaler) convert(node Node, p Path) JNode {
	switch node := node.(type) {
	case *Value:
		return m.newValue(node, p)
	case *Object:
		return m.newObject(node, p)
	case *Pair:
		return m.newPair(node, p)
	case *Array:
		return m.newArray(node, p)
	default:
		return nil
	}
}

func (Marshaler) newPos(p *lexer.Position) *JPos {
	return &JPos{
		Line:   p.Line,
		Column: p.Column,
		Offset: p.Offset,
	}
}

func (m Marshaler) newValue(v *Value, p Path) *JValue {
	if v == nil {
		return nil
	}
	return &JValue{
		Path:   p,
		Pos:    m.newPos(&v.Pos),
		Object: m.newObject(v.Object, p),
		Array:  m.newArray(v.Array, p),
		String: v.String,
		Number: v.Number,
		True:   v.True,
		False:  v.False,
		Null:   v.Null,
	}
}

func (m Marshaler) newObject(v *Object, p Path) *JObject {
	if v == nil || len(v.Pairs) == 0 {
		return nil
	}
	pairs := make([]*JPair, len(v.Pairs))
	for i, pair := range v.Pairs {
		pairs[i] = m.newPair(pair, p.Add(NewIndexPath(pair.Key)))
	}
	return &JObject{
		Path:  p,
		Pos:   m.newPos(&v.Pos),
		Pairs: pairs,
	}
}

func (m Marshaler) newPair(v *Pair, p Path) *JPair {
	if v == nil {
		return nil
	}
	return &JPair{
		Path:  p,
		Pos:   m.newPos(&v.Pos),
		Key:   v.Key,
		Value: m.newValue(v.Value, p),
	}
}

func (m Marshaler) newArray(v *Array, p Path) *JArray {
	if v == nil || len(v.Items) == 0 {
		return nil
	}
	items := make([]*JValue, len(v.Items))
	for i, x := range v.Items {
		items[i] = m.newValue(x, p.Add(NewIndexPath(i)))
	}
	return &JArray{
		Path:  p,
		Pos:   m.newPos(&v.Pos),
		Items: items,
	}
}

type JPos struct {
	Line   int `json:"line"`
	Column int `json:"column"`
	Offset int `json:"offset"`
}

//go:generate go run github.com/berquerant/marker -method IsJNode -type JValue,JObject,JPair,JArray -output ast_jnode_marker_generated.go

// JNode is Node with additional information added.
// The result of json.Marshal is more organized than Node.
type JNode interface {
	IsJNode()
}

type JValue struct {
	Path   Path     `json:"path"`
	Pos    *JPos    `json:"pos"`
	Object *JObject `json:"object,omitempty"`
	Array  *JArray  `json:"array,omitempty"`
	String *string  `json:"string,omitempty"`
	Number *string  `json:"number,omitempty"`
	True   *string  `json:"true,omitempty"`
	False  *string  `json:"false,omitempty"`
	Null   *string  `json:"null,omitempty"`
}

type JObject struct {
	Path  Path     `json:"path"`
	Pos   *JPos    `json:"pos"`
	Pairs []*JPair `json:"pairs,omitempty"`
}

type JPair struct {
	Path  Path    `json:"path"`
	Pos   *JPos   `json:"pos"`
	Key   string  `json:"key,omitempty"`
	Value *JValue `json:"value,omitempty"`
}

type JArray struct {
	Path  Path      `json:"path"`
	Pos   *JPos     `json:"pos"`
	Items []*JValue `json:"items,omitempty"`
}
