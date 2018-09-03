package arrayandstring

import "testing"

func Test_setZeroes(t *testing.T) {
	mat1 := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	want1 := [][]int{
		[]int{1, 0, 1},
		[]int{0, 0, 0},
		[]int{1, 0, 1},
	}
	mat2 := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	want2 := [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 4, 5, 0},
		[]int{0, 3, 1, 0},
	}

	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"0", args{mat1}, want1},
		{"1", args{mat2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
		})
	}
}
