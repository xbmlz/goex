package stream

import (
	"runtime"
	"sync"
)

// processItem 处理单个元素，应用所有操作
func processItem[T any](item T, operations []Operation[T]) (T, bool) {
	res := item
	ok := true
	for _, op := range operations {
		res, ok = op.Apply(res)
		if !ok {
			break
		}
	}
	return res, ok
}

func CollectParallel[T any](s *Stream[T]) []T {
	if !s.ordered {
		// 乱序收集
		var wg sync.WaitGroup
		ch := make(chan T, len(s.items))

		for _, item := range s.items {
			wg.Add(1)
			go func(v T) {
				defer wg.Done()
				res, ok := processItem(v, s.operations)
				if ok {
					ch <- res
				}
			}(item)
		}
		wg.Wait()
		close(ch)

		result := make([]T, 0, len(s.items))
		for v := range ch {
			result = append(result, v)
		}
		return result
	}
	// 保存处理结果和是否保留的信息
	results := make([]T, len(s.items))
	keep := make([]bool, len(s.items))
	var wg sync.WaitGroup

	// 确定工作协程数量 - 不超过CPU核心数或元素数量
	numWorkers := min(len(s.items), runtime.NumCPU())
	jobs := make(chan int, len(s.items))

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				// 处理单个元素
				res, ok := processItem(s.items[idx], s.operations)
				// 保存结果并标记为保留
				if ok {
					results[idx] = res
					keep[idx] = true
				}
			}
		}()
	}

	// 发送任务
	for i := range s.items {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// 收集保留的结果
	finalResult := make([]T, 0, len(s.items))
	for i, ok := range keep {
		if ok {
			finalResult = append(finalResult, results[i])
		}
	}
	return finalResult
}

func (s *Stream[T]) Collect() []T {
	var result []T
	if !s.parallel {
		for _, item := range s.items {
			res := item
			ok := true
			for _, op := range s.operations {
				res, ok = op.Apply(res)
				if !ok {
					break
				}
			}
			if ok {
				result = append(result, res)
			}
		}
	} else {
		result = CollectParallel(s)
	}
	return result
}
