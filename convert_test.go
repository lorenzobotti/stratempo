package stratempo

import (
	"testing"

	"github.com/matryer/is"
)

func TestTokenise(t *testing.T) {
	is := is.New(t)
	testCase := "un milione settecentoquarantuno mila centoventidue"
	expected := []token{
		{typeNumber, 1},
		{typeMagnitude, powInt(10, 6)},
		{typeNumber, 7},
		{typeMagnitude, 100},
		{typeNumber, 40},
		{typeNumber, 1},
		{typeMagnitude, powInt(10, 3)},
		{typeNumber, 100},
		{typeNumber, 20},
		{typeNumber, 2},
		{typeEof, 0},
	}

	converter := newConverter(testCase)
	got, err := converter.tokens()
	is.NoErr(err)

	for i, tok := range got {
		is.True(i < len(expected))
		is.Equal(expected[i], tok)
	}
}

func TestConvert(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"due", 2},
		{"dieci", 10},
		{"ottocento", 800},
		{"cinquecentoventitrè", 523},
		{"centosette", 107},
		{"milleduecentoquaranta", 1240},
		{"novemila settantatre", 9073},
		{"settantasette", 77},
		{"un milione e novecento", 1000900},
		{"un milione e novecento mila quattrocentosettantasette", 1900477},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			is := is.New(t)

			got, err := Convert(tc.input)

			is.NoErr(err)
			is.Equal(tc.expected, got)
		})
	}
}
