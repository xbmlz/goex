package stream

import (
	"testing"
)

func BenchmarkSequentialFilter(b *testing.B) {
	// 构造大数据
	data := make([]int, 1_000_000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Filter(func(x int) bool { return x%2 == 0 }).
			Collect()
	}
}

func BenchmarkParallelFilterOrdered(b *testing.B) {
	data := make([]int, 1_000_000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Parallel().
			Filter(func(x int) bool { return x%2 == 0 }).
			Collect()
	}
}

func BenchmarkParallelFilterUnordered(b *testing.B) {
	data := make([]int, 1_000_000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Parallel().
			Unordered().
			Filter(func(x int) bool { return x%2 == 0 }).
			Collect()
	}
}
