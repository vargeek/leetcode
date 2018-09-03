package linkedlist

import "testing"

func Test_isPalindrome(t *testing.T) {
	l1 := NewListFromInts(1, 2)
	l2 := NewListFromInts(1, 2, 2, 1)
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{l1}, false},
		{"2", args{l2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
