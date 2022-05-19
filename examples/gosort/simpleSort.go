package main

import (
	"fmt"
	"math"
)

func main() {
	//bubbleSort()
	insertionSort()
}

func bubbleSort() {
	someArray := []int{25, 59, 22, 40, 38, 30, 36, 33, 57, 46, 7, 55, 45, 50, 42, 44, 2, 26, 9, 17}
	for j := 0; j < len(someArray); j++ {
		breakLoop := true
		for i := 0; i < len(someArray); i++ {
			if i != len(someArray)-1 {
				if someArray[i] > someArray[i+1] {
					someArray[i], someArray[i+1] = someArray[i+1], someArray[i]
					breakLoop = false
				}
			}
		}
		fmt.Println(j)
		fmt.Println(someArray)
		if breakLoop {
			break
		}
	}
}

func insertionSort() {
	someArray := []int{59, 25, 22, 40, 38, 30, 36, 33, 57, 46, 7, 55, 45, 50, 42, 44, 2, 26, 9, 17}

	for i := 0; i < len(someArray); i++ {
		num := someArray[i]
		hole := i

		for hole > 0 && someArray[hole-1] > num {
			someArray[hole] = someArray[hole-1]
			hole = hole - 1
		}
		someArray[hole] = num
	}

	fmt.Println(someArray)
}

func selectionSort(){
	someArray := []int{}
	sortedArray := []int{}

	//findSmallest
	smallest :+
	for i := 0; i < len(someArray); i++ {
	}
		
	}
}