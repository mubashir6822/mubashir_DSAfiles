package main

import "fmt"

func BinarySearch(detalist []int, key int) int {
	left := 0
	right := len(detalist) - 1
	for left <= right {
		mid := (left + right) / 2
		if detalist[mid] == key {
			return mid
		} else if detalist[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}
func main() {
	detalist := []int{10, 5, 7, 8, 3, 2, 1}
	fmt.Println(BinarySearch(detalist, 7))
	fmt.Println(BinarySearch(detalist, 7))

}



// Binary Search Iterative

// package main
// import "fmt"

// func binarySearch(data int, A []int) bool {
//     low := 0
//     high := len(A) - 1
//     for low <= high{
//         mid := (low + high) / 2
//         if A[mid] < data {
//             low = mid + 1
//         }else{
//             high = mid - 1
//         }
//     }
//     if low == len(A) || A[low] != data {
//         return false
//     }
//     return true
// }

// func main(){
//     items := []int{1,2, 9, 10, 11, 45, 73, 80, 200} // should be a sorted array
//     fmt.Println(binarySearch(63, items))
// }