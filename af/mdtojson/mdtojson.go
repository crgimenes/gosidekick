package mdtojson

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/crgimenes/lex"
)

func Parse(value io.Reader) (ret json.RawMessage, err error) {
	var l lex.Lexer

	l.TokenParsers = []lex.TokenFunction{
		lex.NewLine,
		lex.NotImplemented,
	}

	err = l.Parse(value)
	if err != nil {
		if err != io.EOF {
			return
		}
	}

	for _, t := range l.Tokens {
		fmt.Printf("%v\t%q\n", t.Type, t.Literal)
	}

	return
}
