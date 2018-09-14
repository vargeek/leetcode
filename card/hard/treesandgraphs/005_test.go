package treesandgraphs

import "testing"

func Test_findCircleNum(t *testing.T) {
	m1 := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	want1 := 2

	m2 := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 1},
	}
	want2 := 1

	m3 := [][]int{
		[]int{1, 1, 1, 0, 1, 0, 0},
		[]int{1, 1, 1, 0, 0, 0, 0},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{0, 0, 0, 1, 0, 1, 0},
		[]int{1, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 1, 0},
		[]int{0, 0, 1, 0, 0, 0, 1},
	}
	want3 := 2

	m4 := [][]int{
		[]int{1, 0, 0, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 1},
		[]int{1, 0, 1, 1},
	}
	want4 := 1
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{m1}, want1},
		{"2", args{m2}, want2},
		{"3", args{m3}, want3},
		{"4", args{m4}, want4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
