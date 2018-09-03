package sortingandsearching

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	input1 := []Interval{
		Interval{1, 3},
		Interval{2, 6},
		Interval{8, 10},
		Interval{15, 18},
	}
	want1 := []Interval{
		Interval{1, 6},
		Interval{8, 10},
		Interval{15, 18},
	}
	input2 := []Interval{
		Interval{1, 4},
		Interval{4, 5},
	}
	want2 := []Interval{
		Interval{1, 5},
	}

	input3 := []Interval{}
	want3 := []Interval{}

	input4 := []Interval{
		Interval{0, 5},
	}
	want4 := []Interval{
		Interval{0, 5},
	}

	input5 := []Interval{
		Interval{0, 5},
		Interval{6, 10},
	}
	want5 := []Interval{
		Interval{0, 5},
		Interval{6, 10},
	}

	input6 := []Interval{
		Interval{6, 10},
		Interval{0, 5},
	}
	want6 := []Interval{
		Interval{0, 5},
		Interval{6, 10},
	}

	input7 := []Interval{
		Interval{0, 5},
		Interval{4, 10},
	}
	want7 := []Interval{
		Interval{0, 10},
	}

	type args struct {
		intervals []Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"1", args{input1}, want1},
		{"2", args{input2}, want2},
		{"3", args{input3}, want3},
		{"4", args{input4}, want4},
		{"5", args{input5}, want5},
		{"6", args{input6}, want6},
		{"7", args{input7}, want7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
