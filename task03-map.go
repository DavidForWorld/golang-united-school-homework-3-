package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	Input := map[string]int{10: "aa", 0: "bb", 500: "cc"}

	keys := make([]string, 0, len(Input))

	for key := range Input {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return Input[keys[i]] < Input[keys[j]]
	})
	return
}
