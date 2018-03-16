package problems

import "testing"
import "fmt"

func Test_CountOne(t *testing.T) {
	if CountOne(1) == 1 {
		t.Fatal("1:", 1)
	}
	if CountOne(2) == 1 {
		t.Log("2:", 2)
	}
	if CountOne(11) == 3 {
		t.Log("11:", 3)
	}
}

func Test_Fn(t *testing.T) {
	for i := 0; i < 200000; i++ {
		if i == Fn(i) {
			fmt.Println(i)
		}
	}
}
