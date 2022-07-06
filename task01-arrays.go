package homework

import "fmt"

func average(input [15]float32) (result float32) {
	array := []int{1, 2, 3, 4, 5, 6}
	N := 6
	sum := 0
	for i := 0; i < N; i++ {
		sum += (array[i])
	}
	avg := (float32(sum)) / (float32(N))
	fmt.Println("Sum = ", sum, "\nAverage = ", avg)
	return
}
