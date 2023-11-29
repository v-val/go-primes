package main

import (
	"github.com/jessevdk/go-flags"
	L "github.com/sirupsen/logrus"
	"os"
)

func init() {
	L.SetLevel(L.TraceLevel)
	L.Trace("Enabled")
}

type Options struct {
	NPrimes    int    `short:"n" description:"Number of primes to find"`
	ReadFile   string `short:"r" description:"Read already found primes from file"`
	NoCompress bool   `short:"C" description:"Do not compress output file"`
	TestOnly   bool   `short:"T" description:"Run only test then exit"`
}

func main() {

	var opts Options = Options{
		NPrimes:  1_000_000,
		TestOnly: false,
	}

	if args, err := flags.ParseArgs(&opts, os.Args); err != nil {
		L.Panicf(`%v (left args %v)`, err, args)
	} else {
		L.Debugf(`Options: %v, left args %v`, opts, args)
	}

	if opts.TestOnly {

	}

	if opts.ReadFile != "" {
		primes, err := ReadPrimesDump(opts.ReadFile)
		if err != nil {
			L.Panic(err)
		}
		L.Infof("read %d primes %s", len(primes), P(primes))
	} else {
		primes := make([]prime_value_type, 0, opts.NPrimes)
		primes = MakePrimes(primes, opts.NPrimes)
		DumpPrimes(primes, !opts.NoCompress)
	}
}
