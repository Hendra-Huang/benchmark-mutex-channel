package cachemutex

import (
	"strconv"
	"testing"
)

func benchmarkMutexAdd(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			c.Add(strconv.Itoa(i), T(i))
		}
	}
}

func BenchmarkMutexAdd10(b *testing.B)   { benchmarkMutexAdd(10, b) }
func BenchmarkMutexAdd1000(b *testing.B) { benchmarkMutexAdd(1000, b) }

func benchmarkMutexAddConcurrent(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go func(i int) {
				c.Add(strconv.Itoa(i), T(i))
			}(i)
		}
	}
}

func BenchmarkMutexAddConcurrent1000(b *testing.B) { benchmarkMutexAddConcurrent(1000, b) }

func benchmarkMutexGet(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			c.Get(strconv.Itoa(i))
		}
	}
}

func BenchmarkMutexGet10(b *testing.B)   { benchmarkMutexGet(10, b) }
func BenchmarkMutexGet1000(b *testing.B) { benchmarkMutexGet(1000, b) }

func benchmarkMutexGetConcurrent(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go func(i int) {
				c.Get(strconv.Itoa(i))
			}(i)
		}
	}
}

func BenchmarkMutexGetConcurrent1000(b *testing.B) { benchmarkMutexGetConcurrent(1000, b) }
