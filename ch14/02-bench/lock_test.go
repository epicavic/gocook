package bench

import (
	"sync"
	"testing"
)

func BenchmarkCounterAdd(b *testing.B) {
	c := Counter{0, &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Add(1)
	}
}

func BenchmarkCounterRead(b *testing.B) {
	c := Counter{0, &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Read()
	}
}

func BenchmarkCounterAddRead(b *testing.B) {
	c := Counter{0, &sync.RWMutex{}}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(1)
			c.Read()
		}
	})
}
