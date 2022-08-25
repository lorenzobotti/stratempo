package stratempo

import "testing"

func BenchmarkConvert(b *testing.B) {
	testCase := "un miliardo settecento milioni e novecento mila quattrocentosettantasette"
	expected := 1700900477

	for i := 0; i < b.N; i++ {
		got, err := Convert(testCase)
		if err != nil {
			b.Fail()
		}

		if got != expected {
			b.Fail()
		}
	}
}
