package hard

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	input1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	want1 := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	input2 := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	want2 := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	input3 := [][]int{
		[]int{1, 2, 3, 4},
	}
	want3 := []int{1, 2, 3, 4}

	input4 := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{4},
	}
	want4 := []int{1, 2, 3, 4}

	input5 := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	want5 := []int{1, 2, 4, 3}

	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{input1}, want1},
		{"2", args{input2}, want2},
		{"3", args{input3}, want3},
		{"4", args{input4}, want4},
		{"5", args{input5}, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
