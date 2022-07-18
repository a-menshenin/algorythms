package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

func main() {
	// https://habr.com/ru/post/347378/
	// алгоритм, вычисляющий максимумы всех подпоследовательностей
	// длины k в числовой последовательности длины n
	data := []int{2, 7, 6, 4, 1, 5, 8}
	n := len(data)
	k := 4
	fmt.Print(range_maximums(data, n, k))
}

func insert_maximum_candidate(v int, maximums *deque.Deque[int]) {
	for {
		if maximums.Len() > 0 && maximums.Back() < v {
			maximums.PopBack()
		} else {
			break
		}
	}
	maximums.PushBack(v)
}

func remove_maximum_candidate(v int, maximums *deque.Deque[int]) {
	for {
		if maximums.Len() > 0 && maximums.Front() == v {
			maximums.PopFront()
		} else {
			break
		}
	}
}

func range_maximums(data []int, n, k int) []int {
	maximums := &deque.Deque[int]{}
	maximums.SetMinCapacity(uint(k))
	for i := 0; i < k; i++ {
		insert_maximum_candidate(data[i], maximums)
	}
	result := make([]int, n-k+1)
	for i := k - 1; i < n; i++ {
		insert_maximum_candidate(data[i], maximums)
		result[i-k+1] = maximums.Front()
		remove_maximum_candidate(data[i-k+1], maximums)
	}
	return result
}
