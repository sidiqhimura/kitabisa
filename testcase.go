package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Println("sum :", sum(1, 2))
	fmt.Println("multiply :", multiply(1, 2))
	fmt.Println("prime :", firstPrime(10))
	fmt.Println("fibonacci :", firstFibonacci(10))
}

func sum(x, y int) int {
	return (x + y)
}

func multiply(x, y int) int {
	return (x * y)
}

func firstPrime(n int) []int64 {
	var prime []int64
	var p int = 0
	var i int64 = 1

	for p < n {
		if big.NewInt(i).ProbablyPrime(0) {
			prime = append(prime, i)
			p++
		}
		i++
	}
	return prime
}

func firstFibonacci(n int) []int {
	var seq = make([]int, 2)
	seq[0] = 0
	seq[1] = 1
	for i := 2; i < n; i++ {
		x := (seq[i-2] + seq[i-1])
		seq = append(seq, x)
	}
	return seq
}
