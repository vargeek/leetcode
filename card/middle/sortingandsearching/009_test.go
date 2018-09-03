package sortingandsearching

import "testing"

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	target1 := 5
	want1 := true

	target2 := 20
	want2 := false

	matrix3 := [][]int{}
	target3 := 0
	want3 := false

	matrix4 := [][]int{
		[]int{1},
	}
	target4 := 1
	want4 := true

	matrix5 := [][]int{
		[]int{1},
	}
	target5 := 2
	want5 := false

	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{matrix, target1}, want1},
		{"2", args{matrix, target2}, want2},
		{"3", args{matrix3, target3}, want3},
		{"4", args{matrix4, target4}, want4},
		{"5", args{matrix5, target5}, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
