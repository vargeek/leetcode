package dp

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 5}, 11}, 3},
		{"2", args{[]int{2}, 3}, -1},
		{"3", args{[]int{1, 3, 5}, 8}, 2},
		{"4", args{[]int{245, 331, 498, 374, 314, 407, 494, 285, 4, 435}, 3725}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
