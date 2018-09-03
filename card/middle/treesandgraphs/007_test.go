package treesandgraphs

import "testing"

func Test_numIslands(t *testing.T) {
	grid1 := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	want1 := 1

	grid2 := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	want2 := 3

	grid3 := [][]byte{
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
	}
	want3 := 1

	grid4 := [][]byte{
		[]byte{'0', '0', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	want4 := 0

	grid5 := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	want5 := 1

	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{grid1}, want1},
		{"2", args{grid2}, want2},
		{"3", args{grid3}, want3},
		{"4", args{grid4}, want4},
		{"5", args{grid5}, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
