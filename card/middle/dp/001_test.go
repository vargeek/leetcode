package dp

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"2", args{[]int{3, 2, 1, 0, 4}}, false},
		{"3", args{[]int{0, 1, 2, 3, 4, 5}}, false},
		{"4", args{[]int{1, 2, 3, 4, 5}}, true},
		{"5", args{[]int{1, 1, 1, 1, 1, 1}}, true},
		{"6", args{[]int{1, 1, 1, 0, 1, 1, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
