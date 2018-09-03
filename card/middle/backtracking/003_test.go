package backtracking

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	input1 := []int{1, 2, 3}
	want1 := [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 2, 1},
		[]int{3, 1, 2},
	}
	input2 := []int{}
	want2 := [][]int{}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{input1}, want1},
		{"2", args{input2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
