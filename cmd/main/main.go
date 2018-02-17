package main

import (
	"fmt"

	"github.com/515hikaru/prime-go/eratosthenes"
)

func main() {
	primes := eratosthenes.PrimeSeqence(100)
	for i := 0; i < len(primes); i++ {
		tmp := primes[i]
		if tmp != 0 {
			fmt.Println(primes[i])
		}
	}
}
