package main

import (
	"fmt"
	"sync"

	"main/atomic"
)

func main() {
	o := atomic.NewOrdinal()
	m := atomic.NewSafeMap()
	o.Init(1123)
	fmt.Println("initial ordinal is:", o.GetOrdinal())
	fmt.Printf("initial map is:%+v\n", m)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		fmt.Println("Setting map key/value and incrementing ordinal number: ", i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(i, fmt.Sprintf("i=%d", i))
			o.Increment()
		}(i)
	}
	wg.Wait()

	// check if all map values are set
	for i := 0; i < 10; i++ {
		if _, err := m.Get(i); err != nil {
			panic(err)
		}
	}
	fmt.Println("final ordinal is:", o.GetOrdinal())
	fmt.Printf("final map is:%+v\n", m)
}
