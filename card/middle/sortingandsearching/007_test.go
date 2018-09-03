package sortingandsearching

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"3", args{[]int{4, 5, 6, 7, 0, 1, 2}, 4}, 0},
		{"4", args{[]int{4, 5, 6, 7, 0, 1, 2}, 2}, 6},
		{"5", args{[]int{4, 5, 6, 7, 0, 1, 2}, 7}, 3},
		{"6", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchPivot(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 1, 2, 3, 4, 5, 6, 7}}, 7},
		{"2", args{[]int{7, 0, 1, 2, 3, 4, 5, 6}}, 0},
		{"3", args{[]int{6, 7, 0, 1, 2, 3, 4, 5}}, 1},
		{"4", args{[]int{5, 6, 7, 0, 1, 2, 3, 4}}, 2},
		{"5", args{[]int{4, 5, 6, 7, 0, 1, 2, 3}}, 3},
		{"6", args{[]int{3, 4, 5, 6, 7, 0, 1, 2}}, 4},
		{"7", args{[]int{2, 3, 4, 5, 6, 7, 0, 1}}, 5},
		{"8", args{[]int{1, 2, 3, 4, 5, 6, 7, 0}}, 6},

		{"9", args{[]int{6, 0, 1, 2, 3, 4, 5}}, 0},
		{"10", args{[]int{5, 6, 0, 1, 2, 3, 4}}, 1},
		{"11", args{[]int{4, 5, 6, 0, 1, 2, 3}}, 2},
		{"12", args{[]int{3, 4, 5, 6, 0, 1, 2}}, 3},
		{"13", args{[]int{2, 3, 4, 5, 6, 0, 1}}, 4},
		{"14", args{[]int{1, 2, 3, 4, 5, 6, 0}}, 5},
		{"15", args{[]int{0, 1, 2, 3, 4, 5, 6}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchPivot(tt.args.nums); got != tt.want {
				t.Errorf("searchPivot() = %v, want %v", got, tt.want)
			}
		})
	}
}
