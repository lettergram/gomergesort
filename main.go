package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {

  size := 50

  slice := make([]int, size, size)
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < size; i++ {
      r := rand.Intn(99999) - rand.Intn(99999)
      slice[i] = r
  }

  fmt.Println("----- unsorted ------")
  fmt.Println(slice)
  fmt.Println("----- sorted -----")
  fmt.Println(MergeSort(slice))
  
}

// Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int{

	if len(slice) < 2{
		return slice
	}
	
	mid := (len(slice))/2
  left  := MergeSort(slice[:mid])
  right := MergeSort(slice[mid:])

	return Merge(left, right)
}

// Merges left and right slice into newly created slice
func Merge(left, right []int)[]int{
	
	size := len(left) + len(right)
	slice := make([]int, size, size)
	
	i := 0
	j := 0

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
