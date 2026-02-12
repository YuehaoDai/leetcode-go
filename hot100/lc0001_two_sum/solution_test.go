package lc0001_two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case3", []int{3, 3}, 6, []int{0, 1}},
		{"nil", []int{}, 5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if len(got) != len(tt.want) {
				t.Fatalf("len mismatch: got=%v want=%v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("mismatch: got=%v want=%v", got, tt.want)
				}
			}
		})
	}
}
