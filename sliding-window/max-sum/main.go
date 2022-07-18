package main

import "fmt"

func main() {
	// https://habr.com/ru/post/347378/
	data := []int{2, 5, -8, 3, 9, -2, -4, 0, 2, 6, -3, -5, 0, -4, 2, 3, 1}
	fmt.Printf("[bestStart, bestLength, maxSum] = %d", max_sum_subsequence(data, len(data)))
}

func max_sum_subsequence(data []int, n int) []int {
	bestStart := 0
	bestLength := 1
	maxSum := data[0]
	curStart := bestStart
	curLength := bestLength
	curSum := maxSum
	for i := 1; i < n; i++ {
		if curSum < 0 {
			curStart = i
			curLength = 1
			curSum = data[i]
		} else {
			curLength += 1
			curSum += data[i]
		}
		if curSum > maxSum {
			bestStart = curStart
			bestLength = curLength
			maxSum = curSum
		}
	}
	return []int{bestStart, bestLength, maxSum}
}
