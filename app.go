package main

import (
	"github.com/jessevdk/go-flags"
	L "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os"
	"path/filepath"
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

func GetTemp() string {
	r := os.Getenv("V_TEMP")
	if len(r) == 0 {
		r, _ = os.UserHomeDir()
		r = filepath.Join(r, "temp")
	}
	return r
}

func init() {
	L.SetLevel(L.TraceLevel)
	L.Trace("Enabled")
}

type Options struct {
	NPrimes int `short:"n" description:"Number of primes to find"`
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

type prime_value_type uint64

var primes []prime_value_type

func is_prime(n prime_value_type) bool {
	q := prime_value_type(math.Floor(math.Sqrt(float64(n))))
	for i := 0; primes[i] <= q; i++ {
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}

func make_primes(nprimes int) {

	primes = make([]prime_value_type, 0, nprimes)
	primes = append(primes, 2, 3)

	L.Infof("Start search of %d primes_big", nprimes)
	startTime := time.Now()

	for n := primes[len(primes)-1]; len(primes) < nprimes; n += 2 {
		if is_prime(n) {
			primes = append(primes, n)
		}
	}

	finishTime := time.Now()
	jobDuration := finishTime.Sub(startTime)
	L.Infof("Done: %d primes found in %v at %.1f/s, last %v (0x%016[4]x)",
		len(primes), jobDuration, float64(nprimes)/jobDuration.Seconds(), primes[len(primes)-1])
}

func main() {
	var opts Options = Options{
		NPrimes: 1_000_000,
	}
	if args, err := flags.ParseArgs(&opts, os.Args); err != nil {
		L.Panicf(`%v (left args %v)`, err, args)
	} else {
		L.Debugf(`Options: %v, left args %v`, opts, args)
	}
	nprimes := opts.NPrimes

	make_primes(nprimes)
}
