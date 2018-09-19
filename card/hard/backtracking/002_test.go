package backtracking

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	words1 := []string{"oath", "pea", "eat", "rain"}
	board1 := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	want1 := []string{"oath", "eat"}

	words2 := []string{"a"}
	board2 := [][]byte{
		[]byte{'a'},
	}
	want2 := []string{"a"}

	words3 := []string{"a", "a"}
	board3 := [][]byte{
		[]byte{'a'},
	}
	want3 := []string{"a"}

	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{board1, words1}, want1},
		{"2", args{board2, words2}, want2},
		{"3", args{board3, words3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
