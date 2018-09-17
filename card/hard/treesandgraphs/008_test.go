package treesandgraphs

import "testing"

func Test_longestIncreasingPath(t *testing.T) {

	nums1 := [][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	}
	want1 := 4

	nums2 := [][]int{
		[]int{3, 4, 5},
		[]int{3, 2, 6},
		[]int{2, 2, 1},
	}
	want2 := 4

	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums1}, want1},
		{"2", args{nums2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
