package _347_top_k_frequent_elements

import "testing"

func TestTopKFrequent(t *testing.T) {
	topKFrequent([]int{1, 1, 1, 2, 2, 3, 3, 3}, 2)
	topKFrequent([]int{1}, 1)
}
