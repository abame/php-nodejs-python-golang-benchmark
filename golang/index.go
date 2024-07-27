
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array[] int)[]int {
   for i:=0; i< len(array)-1; i++ {
      for j:=0; j < len(array)-i-1; j++ {
         if (array[j] > array[j+1]) {
            array[j], array[j+1] = array[j+1], array[j]
         }
      }
   }
   return array
}

func main() {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		array[i] = rand.Intn(10000)
	}

	start := time.Now()
	BubbleSort(array)
	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s\n", elapsed)
}
