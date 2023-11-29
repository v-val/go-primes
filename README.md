# go-primes
Finds prime numbers, reports time taken.
Optionally:
* Dumps found primes to file.
* Loads primes from dump.
* Starts web-service returning 5 random prime numbers.

Usage:
```shell
primes [OPTIONS]

Application Options:
  -n=         Number of primes to find (default: 1000000)
  -r=         Read already found primes from file
  -D          Do not dump results to file
  -C          Do not compress output file
  -T          Run only test then exit
  -p=         Port to listen
```