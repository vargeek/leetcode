package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	n1 := 2
	pres1 := [][]int{
		[]int{1, 0},
	}
	want1 := []int{0, 1}

	n2 := 4
	pres2 := [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{3, 1},
		[]int{3, 2},
	}
	want2 := []int{0, 1, 2, 3}
	// want2 := []int{0,2,1,3}

	n3 := 2
	pres3 := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}
	want3 := []int{}

	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{n1, pres1}, want1},
		{"2", args{n2, pres2}, want2},
		{"3", args{n3, pres3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
