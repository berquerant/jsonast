package jsonast

import (
	"io"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func Parse(r io.Reader) (*JSON, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	root, err := jsonParser.ParseBytes("", b)
	if err != nil {
		return nil, err
	}
	return root, nil
}

type JSON struct {
	Value *Value `parser:"@@" json:"value"`
}

func (j *JSON) MarshalJSON() ([]byte, error) {
	var m Marshaler
	return m.Marshal(j.Value)
}

// https://www.json.org/json-en.html
// https://github.com/alecthomas/participle/blob/master/_examples/json/main.go
var (
	jsonLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "String", Pattern: `"(\\"|[^"])*"`},
		{Name: "Number", Pattern: `[-+]?\d+(\.\d+)?([eE][-+]?\d+)?`},
		{Name: "Punct", Pattern: `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
		{Name: "Null", Pattern: "null"},
		{Name: "True", Pattern: "true"},
		{Name: "False", Pattern: "false"},
		{Name: "Whitespace", Pattern: `[ \t\n\r]+`},
	})

	jsonParser = participle.MustBuild[JSON](
		participle.Lexer(jsonLexer),
		participle.Unquote("String"),
		participle.Elide("Whitespace"),
		participle.UseLookahead(2),
	)
)

//go:generate go run github.com/berquerant/marker -method IsNode -type Value,Object,Pair,Array -output ast_node_marker_generated.go

type Node interface {
	IsNode()
}

type Value struct {
	Pos    lexer.Position
	Object *Object `parser:"@@ |" json:"object,omitempty"`
	Array  *Array  `parser:"@@ |" json:"array,omitempty"`
	String *string `parser:"@String |" json:"string,omitempty"`
	Number *string `parser:"@Number |" json:"number,omitempty"`
	True   *string `parser:"@True |" json:"true,omitempty"`
	False  *string `parser:"@False |" json:"false,omitempty"`
	Null   *string `parser:"@Null" json:"null,omitempty"`
}

type Object struct {
	Pos   lexer.Position
	Pairs []*Pair `parser:"'{' @@ (',' @@)* '}'" json:"pairs,omitempty"`
}

type Pair struct {
	Pos   lexer.Position
	Key   string `parser:"@String ':'" json:"key,omitempty"`
	Value *Value `parser:"@@" json:"value,omitempty"`
}

type Array struct {
	Pos   lexer.Position
	Items []*Value `parser:"'[' @@ (',' @@)* ']'" json:"items,omitempty"`
}

func (v *Value) MarshalJSON() ([]byte, error) {
	var m Marshaler
	return m.Marshal(v)
}

func (v *Object) MarshalJSON() ([]byte, error) {
	var m Marshaler
	return m.Marshal(v)
}

func (v *Pair) MarshalJSON() ([]byte, error) {
	var m Marshaler
	return m.Marshal(v)
}

func (v *Array) MarshalJSON() ([]byte, error) {
	var m Marshaler
	return m.Marshal(v)
}
