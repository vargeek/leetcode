package linkedlist

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	list1 := newListNodeFromInts(4, 2, 1, 3)
	want1 := newListNodeFromInts(1, 2, 3, 4)

	list2 := newListNodeFromInts(-1, 5, 3, 4, 0)
	want2 := newListNodeFromInts(-1, 0, 3, 4, 5)

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{list1}, want1},
		{"2", args{list2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
