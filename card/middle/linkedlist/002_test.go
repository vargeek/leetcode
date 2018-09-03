package linkedlist

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	list1 := newListNodeFromInts(1, 2, 3, 4, 5)
	want1 := newListNodeFromInts(1, 3, 5, 2, 4)
	list2 := newListNodeFromInts(2, 1, 3, 5, 6, 4, 7)
	want2 := newListNodeFromInts(2, 3, 6, 7, 1, 5, 4)

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{list1}, want1},
		{"1", args{list2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
