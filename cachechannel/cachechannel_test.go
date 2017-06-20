package cachechannel

import (
	"strconv"
	"testing"
)

func benchmarkChannelAdd(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			c.Add(strconv.Itoa(i), T(i))
		}
	}
}

func BenchmarkChannelAdd10(b *testing.B)   { benchmarkChannelAdd(10, b) }
func BenchmarkChannelAdd1000(b *testing.B) { benchmarkChannelAdd(1000, b) }

func benchmarkChannelAddConcurrent(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go func(i int) {
				c.Add(strconv.Itoa(i), T(i))
			}(i)
		}
	}
}

func BenchmarkChannelAddConcurrent1000(b *testing.B) { benchmarkChannelAddConcurrent(1000, b) }

func benchmarkChannelGet(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			c.Get(strconv.Itoa(i))
		}
	}
}

func BenchmarkChannelGet10(b *testing.B)   { benchmarkChannelGet(10, b) }
func BenchmarkChannelGet1000(b *testing.B) { benchmarkChannelGet(1000, b) }

func benchmarkChannelGetConcurrent(count int, b *testing.B) {
	c := New()

	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go func(i int) {
				c.Get(strconv.Itoa(i))
			}(i)
		}
	}
}

func BenchmarkChannelGetConcurrent1000(b *testing.B) { benchmarkChannelGetConcurrent(1000, b) }
