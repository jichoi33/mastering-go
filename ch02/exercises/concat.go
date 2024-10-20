package main

import "fmt"

func concatArraysToSlice(arr1, arr2 [2]int) []int {
	return append(arr1[:], arr2[:]...)
}

func concatArraysToArray(arr1, arr2 [2]int) [4]int {
	newArray := [4]int{}
	for i := 0; i < 2; i++ {
		newArray[i] = arr1[i]
	}
	for i := 0; i < 2; i++ {
		newArray[i+2] = arr2[i]
	}
	return newArray
}

func concatSlicesToArray(slice1, slice2 []int) [4]int {
	newArray := [4]int{}
	for i := 0; i < 2; i++ {
		newArray[i] = slice1[i]
	}
	for i := 0; i < 2; i++ {
		newArray[i+2] = slice2[i]
	}
	return newArray
}

func main() {
	arr1 := [2]int{1, 2}
	arr2 := [2]int{3, 4}
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}

	fmt.Println(concatArraysToSlice(arr1, arr2))
	fmt.Println(concatArraysToArray(arr1, arr2))
	fmt.Println(concatSlicesToArray(slice1, slice2))
}
