package object

import (
	"bytes"
	"fmt"
	"strings"

	"karp/ast"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

const (
	BUILTIN_OBJ			= "BUILTIN"
	FUNCTION_OBJ		= "FUNCTION"
	INTEGER_OBJ 		= "INTEGER"
	BOOLEAN_OBJ 		= "BOOLEAN"
	STRING_OBJ			= "STRING"
	NULL_OBJ			= "NULL"
	RETURN_VALUE_OBJ 	= "RETURN_VALUE"
	ARRAY_OBJ			= "ARRAY"
	ERROR_OBJ			= "ERROR"
)

type Array struct {
	Elements []Object
}

type Function struct {
	Parameters 	[]*ast.Identifier
	Body		*ast.BlockStatement
	Env 		*Environment
}

type Error struct {
	Message string
}

type Object interface {
	Type()		ObjectType
	Inspect()	string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type String struct {
	Value string
}

type Null struct {}

type ReturnValue struct {
	Value Object
}

type Builtin struct {
	Fn BuiltinFunction
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ } 

func (e *Error) Inspect() string { return "Error: " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ } 

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value } 

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string { return "builtin function" }

func (arr *Array) Type() ObjectType { return ARRAY_OBJ }
func (arr *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range arr.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
