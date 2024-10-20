package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}

	return normalized
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 || len(arguments) > 2 {
		fmt.Println("Need one argument!")
		return
	}

	values := []float64{}

	n, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < n; i++ {
		val := randomFloat(-10, 10)
		values = append(values, val)
	}
	sort.Float64s(values)

	fmt.Println("Number of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	sum := float64(0)
	for _, val := range values {
		sum = sum + val
	}

	meanValue := sum / float64(len(values))
	fmt.Printf("Mean value: %.5f\n", meanValue)

	// Standard deviation
	// https://en.wikipedia.org/wiki/Standard_deviation
	var squared float64
	for i := 0; i < len(values); i++ {
		squared = squared + math.Pow((values[i]-meanValue), 2)
	}

	standardDeviation := math.Sqrt(squared / float64(len(values)))
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	// Now normalize the list of values
	normalized := normalize(values, meanValue, standardDeviation)
	fmt.Println("Normalized:", normalized)
}
