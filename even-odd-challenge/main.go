package main

import (
	"fmt"
)

func main() {
	arr := []int{30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	counter := 0
	arrLenght := len(arr)
	// for i := 0; i < arrLenght; i++ {
	// 	for j := 0; j < arrLenght-i-1; j++ {
	// 		if arr[j+1] < arr[j] {
	// 			temp := arr[j]
	// 			arr[j] = arr[j+1]
	// 			arr[j+1] = temp
	// 		}
	// 		counter++
	// 	}
	// 	fmt.Println(arr)
	// }
	quickSort(arr, 0, len(arr)-1)

	for _, element := range arr {

		if element%2 == 0 {
			fmt.Println(fmt.Sprintf("%d is even", element))
		} else {
			fmt.Println(fmt.Sprintf("%d is odd", element))
		}
	}
	fmt.Println(fmt.Sprintf(fmt.Sprintf("For this %d positions array, %d comparations were performed.", arrLenght, counter)))
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}
