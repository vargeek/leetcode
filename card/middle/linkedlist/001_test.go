package linkedlist

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	first1 := newListNodeFromInts(2, 4, 3)
	second1 := newListNodeFromInts(5, 6, 4)
	want1 := newListNodeFromInts(7, 0, 8)

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{first1, second1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
