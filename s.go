package main

import (
	"fmt"
)

func P(a []prime_value_type) string {
	if len(a) < 7 {
		return fmt.Sprintf("%v", a)
	} else {
		return fmt.Sprintf("%v, %v, %v, ..., %v, %v", a[0], a[1], a[2], a[len(a)-2], a[len(a)-1])
	}
}
