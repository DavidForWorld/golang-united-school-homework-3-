package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	m := map[string]int{"aa": 10, "bb": 0, "cc": 500}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {

	}
	return
}
