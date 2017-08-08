package mdtojson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/crgimenes/lex"
)

func header(r *bufio.Reader, l *lex.Lexer) (ok bool, err error) {
	var value string
	hIndex := 1
	if l.TokRune == '#' {
		ok = true
		value += string(l.TokRune)
		for {
			var size int
			l.TokRune, size, err = r.ReadRune()
			if err == io.EOF {
				break
			}
			if l.TokRune == '#' {
				hIndex++
			}
			l.Offset += size
			if l.TokRune == '\n' {
				break
			}
			value += string(l.TokRune)
		}
	}
	if ok {
		t := lex.Token{Type: fmt.Sprintf("H%d", hIndex),
			Literal: value,
			Line:    l.CurrentLine,
			Offset:  l.Offset,
			Col:     l.CurrentColumn,
		}
		l.Tokens = append(l.Tokens, t)

		l.TokRune = 0
	}
	return
}

func Parse(value io.Reader) (ret json.RawMessage, err error) {
	var l lex.Lexer

	l.TokenParsers = []lex.TokenFunction{
		header,
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
		if t.Type != "NL" &&
			t.Type != "NOT IMPLEMENTED" {

			fmt.Printf("%v\t%q\n", t.Type, t.Literal)
		}
	}

	return
}
