package main

import "fmt"

func main() {
	input := "aababcbdbdbbdcdabd"
	k := 2
	fmt.Print(findLongestSubstring(input, k))
}

func findLongestSubstring(s string, k int) string {
	begin, end := 0, 0
	window, maxWindow := []string{}, []string{}
	freq := make(map[string]int)
	low, high := 0, 0
	for high < len(s) {
		for highEl := high; highEl < len(s); highEl++ {
			if _, ok := freq[string(s[highEl])]; !ok {
				if len(freq) == k {
					if len(window) > len(maxWindow) {
						maxWindow = window
					}
					freq = make(map[string]int)
				}
			}
			// Первое число каждой итерации всегда прибавляем
			window = append(window, string(s[high]))
			freq[string(s[high])] += 1
		}
	}

	return s[begin : end+1]
}
