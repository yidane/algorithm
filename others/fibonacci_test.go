package others

import (
	"testing"
)

var fibonaccis = map[int]int{
	0: 1,
	1: 1,
	2: 2,
	3: 3,
	4: 5,
	5: 8,
	6: 13,
}

func Test_fibonacci(t *testing.T) {
	for k, v := range fibonaccis {
		if fibonacci(k) != v {
			t.Fatal(k, "!=", v)
		}
	}
}

func Test_fibonacci1(t *testing.T) {
	for k, v := range fibonaccis {
		if fibonacci1(k) != v {
			t.Fatal(k, "!=", v)
		}
	}
}
