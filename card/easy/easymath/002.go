package easymath

import (
	"math"
)

// Count Primes
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/744/
// https://leetcode.com/problems/count-primes/

func countPrimes1(n int) int {
	if n <= 2 {
		return 0
	}
	primes := []int{2}

	for i := 3; i < n; i += 2 {
		isPrime := true
		sqrt := int(math.Sqrt(float64(i)))
		for j := 0; j < len(primes) && primes[j] <= sqrt; j++ {
			if i%primes[j] == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return len(primes)
}

func countPrimes(n int) int {
	notPrimes := make([]bool, n)

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if notPrimes[i] == false {
			for j := i * i; j < n; j += i {
				notPrimes[j] = true
			}
		}
	}
	sum := 0
	for i := 2; i < n; i++ {
		if !notPrimes[i] {
			sum++
		}
	}
	return sum
}
