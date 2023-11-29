package main

import (
	L "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

var ZERO = big.NewInt(0)
var TWO = big.NewInt(2)

func is_prime_big(primes_big []*big.Int, n *big.Int) bool {
	q := big.NewInt(0)
	r := big.NewInt(0)
	q.Sqrt(n)
	for _, p := range primes_big {
		if p.CmpAbs(q) == 1 {
			return true
		}
		if r.Mod(n, p).CmpAbs(ZERO) == 0 {
			return false
		}
	}
	return true
}

func make_primes_big(nprimes int) {
	primes_big := make([]*big.Int, 0, nprimes)
	primes_big = append(primes_big, big.NewInt(2), big.NewInt(3))
	c := cap(primes_big)

	L.Infof("Start search of %d primes_big", nprimes)
	startTime := time.Now()
	for n := big.NewInt(0).Add(primes_big[len(primes_big)-1], TWO); len(primes_big) < nprimes; n.Add(n, TWO) {
		if is_prime_big(primes_big, n) {
			p := big.NewInt(0).Set(n)
			primes_big = append(primes_big, p)
			if c2 := cap(primes_big); c != c2 {
				L.Infof("  Capacity %d -> %d", c, c2)
				c = c2
			}
			//p.Bytes()
		}
	}
	finishTime := time.Now()
	jobDuration := finishTime.Sub(startTime)
	L.Infof("Done: %d primes_big found in %vs at %.1f, last %v", len(primes_big), jobDuration, float64(nprimes)/jobDuration.Seconds(), primes_big[len(primes_big)-1])
}
