package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	go randomNumbers(c)
	a := make(chan int)
	go averageCount(c, a)
	go printAverage(a)
	time.Sleep(10 * time.Second)
	close(c)
	fmt.Println("\nFINISH. Channel 'c' is closed")
}

func randomNumbers(c chan int) {
	for {
		numRand := rand.Intn(100)
		c <- numRand
		fmt.Printf("\nSent random number '%v' to channel 'c'   ", numRand)
		time.Sleep(1 * time.Second)
	}
}

func averageCount(c chan int, a chan int) {
	var randSlice []int
	for {
		num := <-c
		randSlice = append(randSlice, num)
		fmt.Println(randSlice)

		var sum int
		for _, val := range randSlice {
			//fmt.Println(val)
			sum += val
		}
		fmt.Println("\nSum: ", sum)
		average := sum / len(randSlice)
		//for TEST fmt.Printf("\nAverage of random numbers in averageCount function is: %v ", average)
		a <- average
	}
}

func printAverage(a chan int) {
	for {
		result := <-a
		fmt.Printf("\nResult of averageCount function is: %v", result)
	}
}
