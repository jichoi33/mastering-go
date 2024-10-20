package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var MIN = 0
var MAX = 94

func random(min, max int, r *rand.Rand) int {
	return r.Intn(max-min) + min
}

func getString(len int64, r *rand.Rand) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX, r)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {
	var LENGTH int64 = 8

	arguments := os.Args
	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			LENGTH = t
		}
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("Using default values...")
	}

	SEED := time.Now().Unix()
	// rand.Seed(SEED)
	r := rand.New(rand.NewSource(SEED))
	fmt.Println(getString(LENGTH, r))
}
