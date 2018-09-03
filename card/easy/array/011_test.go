package easy

import (
	"reflect"
	"testing"
)

func Test_rotateImage(t *testing.T) {
	img1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	img1R := [][]int{
		[]int{7, 4, 1},
		[]int{8, 5, 2},
		[]int{9, 6, 3},
	}
	img2 := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	img2R := [][]int{
		[]int{15, 13, 2, 5},
		[]int{14, 3, 4, 1},
		[]int{12, 6, 8, 9},
		[]int{16, 7, 10, 11},
	}
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{img1}, img1R},
		{"1", args{img2}, img2R},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateImage(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("isValidSudoku() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
