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
}

func randomNumbers(c chan int) {
	for i := 0; i < 5; i++ {
		numRand := rand.Intn(100)
		c <- numRand
		fmt.Printf("\nSent random number '%v' to channel 'c'", numRand)
		time.Sleep(1 * time.Second)
	}
}

func averageCount(c chan int, a chan int) {
	var randSlice []int
	for i := 0; i < 5; i++ {
		num := <-c
		randSlice = append(randSlice, num)
	}
	var sum int
	for i := 0; i < len(randSlice); i++ {
		for _, val := range randSlice {
			//fmt.Println(val)
			sum += val
		}
		fmt.Println("\nSum: ", sum)
		time.Sleep(2 * time.Second)

		average := sum / len(randSlice)
		//for TEST fmt.Printf("\nAverage of random numbers in averageCount function is: %v ", average)
		a <- average
		time.Sleep(5 * time.Second) //щоб зачекати поки завершить роботу після першого  циклу
	}
}

func printAverage(a chan int) {
	result := <-a
	fmt.Printf("\nResult of averageCount function is: %v", result)
}
