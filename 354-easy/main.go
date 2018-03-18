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

// Solve - Given a number A, find the smallest possible value of B+C, if B*C = A
func Solve(input uint64) uint64 {
	//input + 1 will always be a valid answer so make this the default
	solution := uint64(input + 1)

	factors := factor(input)
	sort.Slice(factors, func(i, j int) bool { return factors[i] < factors[j] })

	for i := 0; i < len(factors)/2; i++ {
		x, y := factors[i], factors[len(factors)-1-i]
		if x+y < solution {
			solution = x + y
		}
	}
	return solution
}

func factor(value uint64) []uint64 {
	factors := make([]uint64, 0)

	sqrt := math.Sqrt(float64(value))
	ceiling := uint64(math.Floor(sqrt))
	for i := uint64(1); i <= ceiling; i++ {
		if value%i == 0 {
			factors = append(factors, i, value/i)
		}
	}
	return factors
}
