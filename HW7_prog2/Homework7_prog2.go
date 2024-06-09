package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := make(chan int)
	p := make(chan int)
	v := make(chan int)
	go randNum(r, p, v)
	go extremum(r, p, v)
	time.Sleep(10 * time.Second)
}

func randNum(r chan int, p chan int, v chan int) {
	for i := 0; i < 5; i++ {
		numRand := rand.Intn(100)
		r <- numRand
		fmt.Printf("\nSent random number '%v' to channel 'c'", numRand)
		time.Sleep(1 * time.Second)
	}
	close(r)
	min := <-p
	fmt.Println("Minimum number is: ", min)
	time.Sleep(1 * time.Second)
	max := <-v
	fmt.Println("Maximum number is: ", max)
}

func extremum(r chan int, p chan int, v chan int) {
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
	p <- min
	close(p)
	//fmt.Println("Min num", min)
	max := randSlice[0]
	for _, num := range randSlice {
		if num > max {
			max = num
		}
	}
	v <- max
	close(v)
	//fmt.Println("Max num", max)
}
