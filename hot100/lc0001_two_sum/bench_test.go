package lc0001_two_sum

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	target := 19998

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = twoSum(nums, target)
	}
}
