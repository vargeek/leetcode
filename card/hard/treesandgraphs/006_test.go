package treesandgraphs

import "testing"

func Test_canFinish(t *testing.T) {
	n1 := 2
	pre1 := [][]int{
		[]int{1, 0},
	}
	want1 := true

	n2 := 2
	pre2 := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}
	want2 := false

	n3 := 2
	pre3 := [][]int{
		[]int{1, 1},
		[]int{0, 1},
	}
	want3 := false

	n4 := 3
	pre4 := [][]int{
		[]int{1, 0},
		[]int{2, 0},
	}
	want4 := true

	n5 := 3
	pre5 := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
	}
	want5 := true

	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{n1, pre1}, want1},
		{"2", args{n2, pre2}, want2},
		{"3", args{n3, pre3}, want3},
		{"4", args{n4, pre4}, want4},
		{"5", args{n5, pre5}, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
