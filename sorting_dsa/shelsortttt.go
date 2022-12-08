package main
import "fmt"
func main() {
    fmt.Println("Enter the size of the array")
 var n int
 fmt.Scan(&n)
 A := make([]int, n, 100)
fmt.Println("Enter elements of the array : ")

 for i := 0; i < n; i++ {
  fmt.Scan(&A[i])
 }
fmt.Println("Sorted array:")
   fmt.Println(shellSort(A, n))

}

func shellSort(A []int, n int) []int {

    for interval := int(n/2); interval > 0; interval /= 2 {
     for j := interval; j < n; j++ {
      for k := j; k >= interval && A[k-interval] > A[k]; k -= interval {
       A[k], A[k-interval] = A[k-interval], A[k]
      }
     }
    }
return A
    
}