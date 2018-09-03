package arrayandstring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{"abcabcbb"}, 3},
		{"1", args{"bbbbb"}, 1},
		{"2", args{"pwwkew"}, 3},
		{"3", args{" "}, 1},
		{"4", args{"cdd"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
