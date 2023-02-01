package main

import "fmt"

func partition(A []int, fist int, last int) int {
	pivot := A[fist]
	i := fist + 1
	j := last
	for i < j {
		for A[i] <= pivot && i < last {
			i++
		}
		for A[j] > pivot && j > fist {
			j--
		}
		if i < j {
			// temp := A[i]
			// A[i] = A[j]
			// A[j] = temp
			A[i], A[j] = A[j], A[i]
		}
	}
	A[fist] = A[j]
	A[j] = pivot

	return j
}

func quicksort(A []int, fist, last int) {
	if fist < last {
		var pivot = partition(A, fist, last)
		quicksort(A, fist, pivot-1)
		quicksort(A, pivot+1, last)
	}
}

func main() {
	A := []int{5, 9, 6, 2, 4, 3, 1}
	n := len(A)
	fmt.Println("orignal array:", A)
	quicksort(A, 0, n-1)
	fmt.Println("sorting array:", A)

}
