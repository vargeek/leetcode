package easy

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1,2,3,4,5,6,7", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"-1,-100,3,99", args{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			for i := 0; i < len(tt.args.nums); i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Fatalf("nums[%d](%d)!=want[%d](%d)", i, tt.args.nums[i], i, tt.want[i])
				}
			}
		})
	}
}

// 1 2 3 4 5 6 7
// 5 6 7 1 2 3 4
// 1 6 7 5 2 3 4

// -1, -100, 3, 99
// 3, 99, -1, -100
// 3 -100 -1 99
