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

// Map操作的基准测试
func BenchmarkSequentialMap(b *testing.B) {
	data := make([]int, 1_000_000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Map(func(x int) int { return x * 2 }).
			Collect()
	}
}

func BenchmarkParallelMapOrdered(b *testing.B) {
	data := make([]int, 1_000_000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Parallel().
			Map(func(x int) int { return x * 2 }).
			Collect()
	}
}

func BenchmarkParallelMapUnordered(b *testing.B) {
	data := make([]int, 1_000_000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Parallel().
			Unordered().
			Map(func(x int) int { return x * 2 }).
			Collect()
	}
}

// 计算密集型操作的基准测试
// 计算斐波那契数列的第n项（递归实现，计算成本高）
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func BenchmarkSequentialIntensiveMap(b *testing.B) {
	// 为了避免计算量过大，使用较小的数据集
	data := make([]int, 1_000)
	for i := 0; i < len(data); i++ {
		data[i] = 20 + i%5 // 计算fibonacci(20)到fibonacci(24)之间的值
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Map(func(x int) int { return fibonacci(x) }).
			Collect()
	}
}

func BenchmarkParallelIntensiveMapOrdered(b *testing.B) {
	data := make([]int, 1_000)
	for i := 0; i < len(data); i++ {
		data[i] = 20 + i%5
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Parallel().
			Map(func(x int) int { return fibonacci(x) }).
			Collect()
	}
}

func BenchmarkParallelIntensiveMapUnordered(b *testing.B) {
	data := make([]int, 1_000)
	for i := 0; i < len(data); i++ {
		data[i] = 20 + i%5
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Of(data).
			Parallel().
			Unordered().
			Map(func(x int) int { return fibonacci(x) }).
			Collect()
	}
}
