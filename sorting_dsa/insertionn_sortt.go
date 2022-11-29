package main

import (
	"fmt"
)
func main() {
	fmt.Println("Inseration Sort")
	arr := []int{10, 15, 2, 5, 3, 11}
	fmt.Println("Before sorting")
	fmt.Println(arr)
	fmt.Println("After sorting")
	fmt.Println(insertionSort(arr))
}
func insertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			swap(j, j-1, array)
		}
	}
	return array
}
func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}
