package backtracking

import "testing"

func Test_exist(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	word1 := "ABCCED"
	want1 := true

	word2 := "SEE"
	want2 := true

	word3 := "ABCB"
	want3 := false

	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{board, word1}, want1},
		{"2", args{board, word2}, want2},
		{"3", args{board, word3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
