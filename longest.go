package stratempo

import (
	"sort"
	"strings"
)

var digitsLp = makeLongestPicker(digitsMap)
var magnitudeLp = makeLongestPicker(magnitudeMap)

type namedElement struct {
	name  string
	value int
}

type longestPicker []namedElement

func makeLongestPicker(m map[string]int) longestPicker {
	elements := make([]namedElement, 0, len(m))
	for name, value := range m {
		elements = append(elements, namedElement{name, value})
	}

	sort.Slice(elements, func(i, j int) bool {
		return len(elements[i].name) > len(elements[j].name)
	})
	return elements
}

func (lp longestPicker) match(name string) (value int, found bool, length int) {
	for _, element := range lp {
		if strings.HasPrefix(name, element.name) {
			return element.value, true, len(element.name)
		}
	}

	return 0, false, 0
}
