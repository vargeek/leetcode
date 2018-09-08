package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	lists1 := []*ListNode{
		newListNodeFromInts(1, 4, 5),
		newListNodeFromInts(1, 3, 4),
		newListNodeFromInts(2, 6),
	}
	want1 := newListNodeFromInts(1, 1, 2, 3, 4, 4, 5, 6)

	lists2 := []*ListNode{}
	want2 := newListNodeFromInts()

	lists3 := []*ListNode{
		newListNodeFromInts(1, 4, 5),
	}
	want3 := newListNodeFromInts(1, 4, 5)

	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{lists1}, want1},
		{"2", args{lists2}, want2},
		{"3", args{lists3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
