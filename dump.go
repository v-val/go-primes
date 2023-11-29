package main

import (
	"encoding/binary"
	"fmt"
	L "github.com/sirupsen/logrus"
	"github.com/ulikunitz/xz/lzma"
	"io"
	"os"
	"path/filepath"
	"time"
)

func GetTemp() string {
	r := os.Getenv("V_TEMP")
	if len(r) == 0 {
		r, _ = os.UserHomeDir()
		r = filepath.Join(r, "temp")
	}
	return r
}

func DumpPrimes(primes []prime_value_type, isXzEnabled bool) int {
	var err error
	var r int

	L.Debugf("Compress: %v", isXzEnabled)

	file := filepath.Join(GetTemp(), "primes.dat")
	if isXzEnabled {
		file += ".lzma2"
	}

	t := time.Now()

	buf := make([]byte, 0, len(primes)*prime_value_size /*unsafe.Sizeof(prime_value_type())*/)
	for _, v := range primes {
		if prime_value_size == 4 {
			buf = binary.LittleEndian.AppendUint32(buf, uint32(v))
		}
	}
	t2 := time.Now()

	if isXzEnabled {
		var h *os.File
		if h, err = os.Create(file); err == nil {
			var xz *lzma.Writer2
			if xz, err = lzma.NewWriter2(h); err == nil {
				defer func() {
					if err := xz.Close(); err != nil {
						L.Error(err)
					}
				}()
				r, err = xz.Write(buf)
				if err == nil && r != len(buf) {
					err = fmt.Errorf("write result: actual %d != expected %d", r, len(buf))
				}
			}
		}
	} else {
		err = os.WriteFile(file, buf, 0640)
	}

	{
		d := time.Now().Sub(t)
		L.Infof(`Binarize %d primes as %d bytes in %v (%fB/s)`, len(primes), len(buf), d, float64(len(buf))/d.Seconds())
	}

	if err != nil {
		L.Errorf(`Fail to dump to "%s": %v`, file, err)
	} else {
		d := time.Now().Sub(t2)
		L.Infof(`Wrote %d primes to "%s" as %dB in %v`, len(primes), file, r, d)
	}

	return r
}

func ReadPrimesDump(file string) ([]prime_value_type, error) {
	var result []prime_value_type
	var resultCapacity int
	var err error
	var in *os.File
	if in, err = os.Open(file); err == nil {
		var xz *lzma.Reader2
		if xz, err = lzma.NewReader2(in); err == nil {
			readTotal := 0
			result = make([]prime_value_type, 0, 1024)
			resultCapacity = cap(result)
			buf := make([]byte, 1024*1024)
			for err == nil && !xz.EOS() {
				var n int
				n, err = xz.Read(buf)
				if n > 0 {
					//L.Tracef("read %dB / %d values", n, n/prime_value_size)
					readTotal += n
					for offset := 0; offset < n; offset += prime_value_size {
						result = append(result, prime_value_type(binary.LittleEndian.Uint32(buf[offset:])))
						if c := cap(result); c != resultCapacity {
							//L.Tracef("%d results, capacity %d", len(result), c)
							resultCapacity = c
						}
					}
				}
			}
			if err == io.EOF {
				err = nil
			}
			L.Debugf("Got %d bytes, %d primes:", readTotal, len(result))
		}
	}
	return result, err
}
