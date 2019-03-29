package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func Test_HelloWorld(t *testing.T) {
	fmt.Println("hello world")
	var step int32 = 1

	for i := 1; i <= 10; i++ {
		go func(i int) {
			v := int32(2*i - 1)
			s := atomic.LoadInt32(&step)
			for s != v {
				s = atomic.LoadInt32(&step)
			}
			fmt.Println(v)
			atomic.StoreInt32(&step, v+1)
		}(i)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			v := int32(2 * i)

			s := atomic.LoadInt32(&step)
			for s != v {
				s = atomic.LoadInt32(&step)
			}
			fmt.Println(v)
			atomic.StoreInt32(&step, v+1)
		}(i)
	}

	time.Sleep(3 * time.Second)

	fmt.Println(step)
}
