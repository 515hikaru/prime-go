package eratosthenes

import (
	"math"
)

// PrimeSeqence return prime numbers array.
func PrimeSeqence(length int) []int {
	initialFlgs := make([]int, length)
	flgs := initFlags(initialFlgs)
	primes := make([]int, length)
	maxIter := math.Sqrt(float64(length))
	for i := 2; i <= int(maxIter); i++ {
		if flgs[i] == 0 {
			continue
		}
		for j := i; j < length; j += i {
			if j == i {
				continue
			}
			flgs[j] = 0
		}
	}
	c := 0
	for i := 0; i < len(flgs); i++ {
		if flgs[i] == 1 {
			primes[c] = i
			c++
		}
	}
	return primes
}

func initFlags(flgs []int) []int {
	for i := 2; i < len(flgs); i++ {
		flgs[i] = 1
	}
	return flgs
}
