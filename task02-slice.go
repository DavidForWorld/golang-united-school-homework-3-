package homework

import (
	"sort"
)

func reverse(input []int64) (result []int64) {
	nums := []int64{1, 2, 5, 15}
	ReverseSlice(nums)
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
	return
}
