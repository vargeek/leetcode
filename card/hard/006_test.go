package hard

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 0}}, 3},
		{"2", args{[]int{3, 4, -1, 1}}, 2},
		{"3", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"4", args{[]int{}}, 1},
		{"5", args{[]int{1, 2}}, 3},
		{"6", args{[]int{10}}, 1},
		{"7", args{[]int{2, 4, 6, 8, 3, 1}}, 5},
		{"8", args{[]int{-1, 1, -2, -3, -4}}, 2},
		{"9", args{[]int{2}}, 1},
		{"10", args{[]int{1, 1}}, 2},
		{"11", args{[]int{1, 5, 3, 3, 3, 3, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
