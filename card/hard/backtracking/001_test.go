package backtracking

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	input1 := "aab"
	want1 := [][]string{
		[]string{"aa", "b"},
		[]string{"a", "a", "b"},
	}

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{input1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
