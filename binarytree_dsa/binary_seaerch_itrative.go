//Binary Search Iterative

package main

import "fmt"

func binarySearch(data int, A []int) bool {
	low := 0
	high := len(A) - 1
	for low <= high {
		mid := (low + high) / 2
		if A[mid] < data {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(A) || A[low] != data {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 9, 10, 11, 45, 73, 80, 200} // should be a sorted array
	fmt.Println(binarySearch(63, items))
	fmt.Println(binarySearch(45, items))
}
