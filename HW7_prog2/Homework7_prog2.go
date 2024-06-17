package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := make(chan int)
	v := make(chan minMax)
	go randNum(r, v)
	go extremum(r, v)
	time.Sleep(10 * time.Second)
}

func randNum(r chan int, v chan minMax) {
	for i := 0; i < 5; i++ {
		numRand := rand.Intn(100)
		r <- numRand
		fmt.Printf("\nSent random number '%v' to channel 'c'", numRand)
		time.Sleep(1 * time.Second)
	}
	close(r)
	results := <-v
	fmt.Println("Minimum number is: ", results.min)
	fmt.Println("Maximum number is: ", results.max)
	time.Sleep(1 * time.Second)
}

type minMax struct {
	min int
	max int
}

func extremum(r chan int, v chan minMax) {
	var randSlice []int
	for i := 0; i < 5; i++ {
		num := <-r
		randSlice = append(randSlice, num)
		fmt.Println(randSlice)
	}

	fmt.Println(randSlice)
	time.Sleep(3 * time.Second)
	min := randSlice[0]
	for _, num := range randSlice {
		if num < min {
			min = num
		}
	}

	max := randSlice[0]
	for _, num := range randSlice {
		if num > max {
			max = num
		}
	}

	results := minMax{
		min: min,
		max: max,
	}
	v <- results
	close(v)
}
