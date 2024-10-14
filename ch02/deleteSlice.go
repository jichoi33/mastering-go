package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	// Delete element at index i
	if i > len(aSlice)-1 || i < 0 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// The ... operator auto expands aSlice[i+1:] so that
	// its elements can be appended to aSlice[:i] one by one
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)

	aSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	// Replace element at index i with last element
	aSlice[i] = aSlice[len(aSlice)-1]
	// Remove last element
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)

	aSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	aSlice = slices.Delete(aSlice, 3, 4)
	fmt.Println("After 3rd deletion:", aSlice)
}
