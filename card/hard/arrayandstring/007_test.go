package hard

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{100, 4, 200, 1, 3, 2}}, 4},
		{"2", args{[]int{0}}, 1},
		{"3", args{[]int{1}}, 1},
		{"4", args{[]int{1, 1}}, 1},
		{"5", args{[]int{2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
