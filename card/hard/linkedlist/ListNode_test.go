package linkedlist

import (
	"reflect"
	"testing"
)

func Test_newListNodeFromInts(t *testing.T) {
	input1 := []int{4, 2, 91, 7, 4}
	input2 := []int{}
	input3 := []int{1}
	input4 := []int{1, 1, 2, 2}
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{input1}, input1},
		{"2", args{input2}, input2},
		{"3", args{input3}, input3},
		{"4", args{input4}, input4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newListNodeFromInts(tt.args.vals...).toInts()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newListNodeFromInts().toInts = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_toInts(t *testing.T) {
	input1 := &ListNode{5, &ListNode{2, &ListNode{29, &ListNode{0, nil}}}}
	want1 := []int{5, 2, 29, 0}

	input2 := &ListNode{0, nil}
	want2 := []int{0}

	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"1", fields{input1.Val, input1.Next}, want1},
		{"2", fields{input2.Val, input2.Next}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := l.toInts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNode.toInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
