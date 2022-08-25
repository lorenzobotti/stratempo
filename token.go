package stratempo

import "fmt"

type tokenType int

// type tokenType string

const (
	typeEof = iota
	typeNumber
	typeMagnitude
)

// const (
// 	typeEof       = "eof"
// 	typeNumber    = "number"
// 	typeMagnitude = "magnitude"
// )

type token struct {
	kind  tokenType
	value int
}

func eofToken() token {
	return token{typeEof, 0}
}

func (t token) String() string {
	switch t.kind {
	case typeMagnitude:
		return fmt.Sprint("m", t.value)
	case typeNumber:
		return fmt.Sprint("v", t.value)
	case typeEof:
		return "eof"
	default:
		panic("unknwon token typemn")
	}
}
