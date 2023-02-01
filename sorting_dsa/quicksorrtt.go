package main

import (
	"fmt"
	"strconv"
)

func quickSort(A []int, low, high int) {
	if low < high {

		// Find the pivot index of an lower bound of an array
		var pivot = partition(A, low, high)

		// Apply Divide and conquer strategy
		// to sort the left side and right side of an array
		// respective to the pivot position

		// Left hand side array
		quickSort(A, low, pivot)

		// Right hand side array
		quickSort(A, pivot+1, high)
	}
}

func partition(A []int, low, high int) int {

	// Pick a lowest bound element as an pivot value
	var pivot = A[low]

	var i = low
	var j = high

	for i < j {

		// Increment the index value of "i"
		// till the left most values should be less than or equal to the pivot value
		for A[i] <= pivot && i < high {
			i++
		}

		// Decrement the index value of "j"
		// till the right most values should be greater than to the pivot value
		for A[j] > pivot && j > low {
			j--
		}

		// Interchange the values of present
		// in the index i and j of an array only if i is less than j
		if i < j {
			var temp = A[i]
			A[i] = A[j]
			A[j] = temp
		}
	}

	// Interchange the element in j's poisition to the lower bound of an array
	// and place the Pivot element to the j's position
	A[low] = A[j]
	A[j] = pivot

	// Finally return the appropriate index position of the pivot element
	return j
}

func main() {

	fmt.Println("Enter the size of the array")
	var n int
	fmt.Scan(&n)
	A := make([]int, n, 100)
	fmt.Println("Enter elements : ")

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	quickSort(A, 0, len(A)-1)
	fmt.Print("After Sorting: ")
	for i := 0; i < n; i++ {
		fmt.Print(strconv.Itoa(A[i]) + " ")
	}
}
