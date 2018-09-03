package mmath

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{2.00000, 10}, 1024.00000},
		{"2", args{2.10000, 3}, 9.26100},
		{"3", args{2.00000, -2}, 0.25000},
		{"4", args{0.0, 1}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); math.Abs(got-tt.want) > 0.00001 {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
