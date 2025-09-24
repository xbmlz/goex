package stream

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}

	// 顺序流
	expected := []int{2, 4, 6}
	result := Of(nums).
		Filter(func(x int) bool { return x%2 == 0 }).
		Collect()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sequential Filter failed. Got %v, want %v", result, expected)
	}

	// 并行保序流
	resultParallel := Of(nums).
		Parallel().
		Filter(func(x int) bool { return x%2 == 0 }).
		Collect()
	if !reflect.DeepEqual(resultParallel, expected) {
		t.Errorf("Parallel Ordered Filter failed. Got %v, want %v", resultParallel, expected)
	}

	// 并行乱序流
	resultUnordered := Of(nums).
		Parallel().
		Unordered().
		Filter(func(x int) bool { return x%2 == 0 }).
		Collect()
	// 乱序流只保证元素存在，不保证顺序
	if len(resultUnordered) != len(expected) {
		t.Errorf("Parallel Unordered Filter length mismatch. Got %v, want %v", resultUnordered, expected)
	}
	for _, v := range expected {
		found := false
		for _, u := range resultUnordered {
			if u == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Parallel Unordered Filter missing element %v", v)
		}
	}
}
