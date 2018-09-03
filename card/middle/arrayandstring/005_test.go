package arrayandstring

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{"babad"}, "bab"},
		{"1", args{"cbbd"}, "bb"},
		{"2", args{"babad"}, "bab"},
		{"3", args{"a"}, "a"},
		{"4", args{""}, ""},
		{"5", args{"bb"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
