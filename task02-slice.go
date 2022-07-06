package homework

import "reflect"

func reverse(s interface{}) {
	var n = reflect.ValueOf(s).Len()
	var swap = reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
