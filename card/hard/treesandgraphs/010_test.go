package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	input1 := []int{5, 2, 6, 1}
	want1 := []int{2, 1, 1, 0}

	input2 := []int{-1, -1}
	want2 := []int{0, 0}

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{input1}, want1},
		{"2", args{input2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
