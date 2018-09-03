package easymath

import (
	"testing"
)

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{10}, 4},
		{"1", args{1}, 0},
		{"2", args{2}, 0},
		{"3", args{3}, 1},
		{"4", args{4}, 2},
		{"5", args{499979}, 41537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
	// countPrimes(499979)
}
