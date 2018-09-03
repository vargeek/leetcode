package mmath

import (
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, 3}, 3},
		{"2", args{7, -3}, -2},
		{"3", args{1, 0}, 1<<31 - 1},
		{"4", args{-1 << 31, -1}, 1<<31 - 1},
		{"5", args{1<<31 - 1, 1}, 1<<31 - 1},
		{"6", args{1<<31 - 1, 1<<31 - 1}, 1},
		{"7", args{-2147483648, -1}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}

}
