package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {

  slice := generateSlice(50)
  fmt.Println("\n --- unsorted --- \n", slice)
  fmt.Println("\n--- sorted ---\n", MergeSort(slice))
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {	

  slice := make([]int, size, size)
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < size; i++ {
      slice[i] = rand.Intn(99999) - rand.Intn(99999)
  }
  return slice
}

// Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int{

	if len(slice) < 2{
		return slice
	}
	mid := (len(slice))/2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right []int)[]int{
	
	size := len(left) + len(right)
	slice := make([]int, size, size)
	i, j := 0

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1{
			slice[k] = right[j]
			j++
		}else if j > len(right)-1 && i <= len(left)-1{
			slice[k] = left[i]
			i++
		}else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		}else{
			slice[k] = right[j]
			j++	
		}
	}
	return slice
}
