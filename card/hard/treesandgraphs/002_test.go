package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	board1 := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}

	want1 := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"1", args{board1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("got %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
