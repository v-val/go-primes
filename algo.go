package main

import (
	L "github.com/sirupsen/logrus"
	"math"
	"time"
)

type prime_value_type uint32

const prime_value_size = 4

func IsPrime(primes []prime_value_type, n prime_value_type) bool {
	q := prime_value_type(math.Floor(math.Sqrt(float64(n))))
	for i := 0; primes[i] <= q; i++ {
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}

func MakePrimes(primes []prime_value_type, nprimes int) []prime_value_type {

	primes = append(primes, 2, 3)

	L.Infof("Start search of %d primes", nprimes)
	startTime := time.Now()

	for n := primes[len(primes)-1]; len(primes) < nprimes; n += 2 {
		if IsPrime(primes, n) {
			primes = append(primes, n)
		}
	}

	finishTime := time.Now()
	jobDuration := finishTime.Sub(startTime)
	L.Infof("Done: %d primes found in %v at %.1f/s, last %v (0x%016[4]x)",
		len(primes), jobDuration, float64(nprimes)/jobDuration.Seconds(), primes[len(primes)-1])
	return primes
}
