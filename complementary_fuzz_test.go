package stratempo

import (
	"testing"

	"github.com/lorenzobotti/nombre"
	"github.com/matryer/is"
)

func FuzzBoth(f *testing.F) {
	f.Add(400)
	f.Add(1)
	f.Add(432)
	f.Add(4398)
	f.Add(8543)

	f.Fuzz(func(t *testing.T, input int) {
		is := is.New(t)

		if input > 99999999999999 || input < 0 {
			t.Skip()
		}
		stringified := nombre.Convert(input)
		translated, err := Convert(stringified)
		is.NoErr(err)
		is.Equal(input, translated)
	})
}

func FuzzConvert(f *testing.F) {
	f.Add("cinquantamila")
	f.Add("harruy styles")
	f.Fuzz(func(t *testing.T, input string) {
		Convert(input)
	})
}
