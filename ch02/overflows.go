package main

import (
	"fmt"
	"math"
)

func main() {
	i := math.MaxInt
	for {
		if i == math.MaxInt {
			break
		}
		i++
	}
	fmt.Println("Max:", i)
	fmt.Println("Max overflow:", i+1)

	i = math.MinInt + 1000
	for {
		if i == math.MinInt {
			break
		}
		i--
	}
	fmt.Println("Min:", i)
	fmt.Println("Min overflow:", i-1)
}
