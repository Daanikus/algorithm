package main

import "fmt"

func main() {
	var x2 = []int{1, 13, 5, 7, 9, 11, 8888, 3333, 227, 313123, 88321, 321}
	x2 = []int{3, 1}
	fmt.Println(insertionSort(x2))

}

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		value := arr[i]

		for j := i - 1; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1], arr[j] = arr[j], value
			} else {
				break
			}

		}
	}
	return arr

}

//
//func sortAsc(arr []int) []int {
//	if len(arr) < 1 {
//		return arr
//	}
//	for j := 1; j < len(arr); j++ {
//		key := arr[j]
//		i := j - 1
//
//		for i > 0 && arr[i] > key {
//			arr[i+1] = arr[i]
//			i = i - 1
//		}
//
//		arr[i+1] = key
//
//	}
//	return arr
//}
//
