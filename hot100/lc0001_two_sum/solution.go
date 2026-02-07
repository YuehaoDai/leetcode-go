package lc0001_two_sum

func twoSum(nums []int, target int) []int {
	// TODO: implement
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := seen[target-num]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
