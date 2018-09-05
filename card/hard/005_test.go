package hard

import (
	"reflect"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	input1 := [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}
	want1 := [][]int{
		[]int{0, 0, 0},
		[]int{1, 0, 1},
		[]int{0, 1, 1},
		[]int{0, 1, 0},
	}
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{input1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("board = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
