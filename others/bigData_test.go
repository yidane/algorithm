package others

import (
	"fmt"
	"testing"
)

func Test_findOneInBigData(t *testing.T) {
	c := make(chan int)
	total := 10000
	v := 6789

	go func() {
		for i := 0; i < total; i++ {
			if i == v {
				continue
			}
			c <- i
		}
		close(c)
	}()

	r := findOneInBigData(total, c)

	if r != v {
		t.Fatal(r, "!=", v)
	}
}

func Test_Mod(t *testing.T) {
	i := 1 << 1
	j := 1 << 2

	fmt.Printf("i<<0	%b\n", 1<<0)
	fmt.Printf("i	0%b\n", i)
	fmt.Printf("j	%b\n", j)
	//与 都为1才为1
	fmt.Printf("i&j	%b\n", (1<<1)&(1<<2))
	//或 有1就为1
	fmt.Printf("i|j %b\n", i|j)
	//异或 不同为1，相同为0
	fmt.Printf("i^j %b\n", i^j)
	fmt.Printf("i|j^i %b\n", i|j^i)
	//非
	fmt.Printf("!i	%b\n", ^i)
}
