package bench

import "testing"

func BenchmarkAtomicCounterAdd(b *testing.B) {
	c := AtomicCounter{0}
	for n := 0; n < b.N; n++ {
		c.Add(1)
	}
}

func BenchmarkAtomicCounterRead(b *testing.B) {
	c := AtomicCounter{0}
	for n := 0; n < b.N; n++ {
		c.Read()
	}
}

func BenchmarkAtomicCounterAddRead(b *testing.B) {
	c := AtomicCounter{0}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(1)
			c.Read()
		}
	})
}
