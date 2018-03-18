package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	input := uint64(1234567891011)
	fmt.Println(Solve(input))
}

func Solve(input uint64) uint64 {
	lowest := uint64(input + 1)
	factors := factor(input)
	for i := 0; i < len(factors)/2; i++ {
		x, y := factors[i], factors[len(factors)-1-i]
		if x+y < lowest {
			lowest = x + y
		}
	}
	return lowest
}

func factor(value uint64) []uint64 {
	factors := make([]uint64, 0)

	sqrt := math.Sqrt(float64(value))
	ceiling := uint64(math.Ceil(sqrt))
	for i := uint64(1); i <= ceiling; i++ {
		if value%i == 0 {
			factors = append(factors, i, value/i)
		}
	}

	sort.Slice(factors, func(i, j int) bool { return factors[i] < factors[j] })
	return factors
}
