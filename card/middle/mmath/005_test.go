package mmath

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{4}, 2},
		{"2", args{8}, 2},
		{"3", args{9}, 3},
		{"4", args{10}, 3},
		{"5", args{11}, 3},
		{"6", args{12}, 3},
		{"7", args{13}, 3},
		{"8", args{14}, 3},
		{"9", args{15}, 3},
		{"10", args{16}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
