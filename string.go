package main

import (
	"fmt"
)

type HeadTail struct {
	values_ []prime_value_type
}

func (t HeadTail) String() string {
	if len(t.values_) < 7 {
		return fmt.Sprintf("%v", t.values_)
	} else {
		return fmt.Sprintf("%d, %d, %d, ..., %d, %d",
			t.values_[0], t.values_[1], t.values_[2], t.values_[len(t.values_)-2], t.values_[len(t.values_)-1])
	}
}
