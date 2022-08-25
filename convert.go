package stratempo

import (
	"fmt"
	"unicode"
)

type converter struct {
	input  []rune
	cursor int

	lastToken  token
	buffer     token
	firstToken bool
}

func (c converter) remaining() string {
	return string(c.input[c.cursor:])
}

func (c converter) getCursor() rune {
	if c.cursor < len(c.input) {
		return c.input[c.cursor]
	} else {
		return ' '
	}
}

func (c *converter) skipWhitespace() {
	for ; unicode.IsSpace(c.getCursor()) && c.cursor < len(c.input); c.cursor += 1 {
	}
}

func (c *converter) nextToken() (token, error) {
	if c.buffer != eofToken() {
		tok := c.buffer
		c.buffer = eofToken()
		return tok, nil
	}

	if c.cursor >= len(c.input) {
		return eofToken(), nil
	}

	c.skipWhitespace()

	digit, isDigit, length := digitsLp.match(c.remaining())
	if isDigit {
		c.cursor += length
		tok := token{typeNumber, digit}

		// "cento" (one hundred) can be a number on its own
		// (ex. "centoventi") or an order of magnitude (ex. "settecento").
		// we differentiate them by checking wether the last token was a
		// number or a magnitude
		if digit == 100 && c.lastToken.kind == typeNumber {
			tok = token{typeMagnitude, 100}
		}
		if digit == 1000 && c.lastToken.kind == typeNumber {
			tok = token{typeMagnitude, 1000}
		}

		c.lastToken = tok
		return tok, nil
	}

	magnitude, isMagnitude, length := magnitudeLp.match(c.remaining())
	if isMagnitude {
		c.cursor += length
		tok := token{typeMagnitude, magnitude}

		if (c.lastToken.kind == typeMagnitude && c.lastToken.value != 100) || c.lastToken.kind == typeEof {
			c.buffer = tok
			return token{typeNumber, 1}, nil
		}

		c.lastToken = tok
		return tok, nil
	}

	// ignore the 'e' when it separates orders of magnitude.
	// ex: mille e novecento
	if c.getCursor() == 'e' {
		c.cursor += 1
		return c.nextToken()
	}

	return eofToken(), fmt.Errorf("wrong word: '%s'", c.remaining())
}

func (c *converter) tokens() ([]token, error) {
	out := []token{}
	for {
		next, err := c.nextToken()
		if err != nil {
			return out, err
		}

		out = append(out, next)

		if next == eofToken() {
			return out, nil
		}
	}
}

func (c *converter) convert() (int, error) {
	totalResult := 0
	currentOrderOfMagnitude := 0
	lastKind := tokenType(typeEof)

	// fmt.Println()
	// fmt.Println(c.remaining())

	for {
		tok, err := c.nextToken()
		if err != nil {
			return 0, err
		}

		// fmt.Println(tok)

		if tok.kind == typeNumber {
			if lastKind == typeMagnitude {
				totalResult += currentOrderOfMagnitude
				currentOrderOfMagnitude = 0
			}
			currentOrderOfMagnitude += tok.value

		} else if tok.kind == typeMagnitude {
			currentOrderOfMagnitude *= tok.value
		} else if tok.kind == typeEof {
			return totalResult + currentOrderOfMagnitude, nil
		}

		lastKind = tok.kind
	}
}

func newConverter(input string) converter {
	return converter{
		input:      []rune(input),
		cursor:     0,
		buffer:     eofToken(),
		lastToken:  eofToken(),
		firstToken: true,
	}
}

func Convert(input string) (int, error) {
	converter := newConverter(input)
	return converter.convert()
}
