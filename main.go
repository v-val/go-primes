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
	NoDump     bool   `short:"D" description:"Do not dump results to file"`
	NoCompress bool   `short:"C" description:"Do not compress output file"`
	TestOnly   bool   `short:"T" description:"Run only test then exit"`
	Port       uint16 `short:"p" description:"Port to listen"`
}

func main() {

	var opts Options = Options{
		NPrimes:  1_000_000,
		TestOnly: false,
		NoDump:   true,
	}

	if args, err := flags.ParseArgs(&opts, os.Args); err != nil {
		L.Panicf(`%v (left args %v)`, err, args)
	} else {
		L.Debugf(`Options: %v, left args %v`, opts, args)
	}

	if opts.TestOnly {

	}

	var primes []prime_value_type

	if opts.ReadFile != "" {
		r, err := ReadPrimesDump(opts.ReadFile)
		if err != nil {
			L.Panic(err)
		}
		L.Infof("loaded %d primes: %s.", len(r), HeadTail{r})
		primes = r
	} else {
		r := make([]prime_value_type, 0, opts.NPrimes)
		r = MakePrimes(r, opts.NPrimes)
		L.Infof("built %d primes: %s.", len(r), HeadTail{r})
		primes = r
		//-r/Volumes/RAMDisk/primes.dat.lzma2
		if !opts.NoDump {
			DumpPrimes(r, !opts.NoCompress)
		}
	}

	if opts.Port != 0 {
		L.Infof("have %d primes: %s.", len(primes), HeadTail{primes})
		hostname, err := os.Hostname()
		if err != nil {
			L.Panic(err)
		}
		L.Tracef("hostname %s", hostname)
		service := &HttpService{
			Id:     hostname,
			Port:   opts.Port,
			Values: primes,
		}
		L.Fatal(service.Run())
	}
}
