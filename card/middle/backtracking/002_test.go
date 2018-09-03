package backtracking

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	n1 := 3
	want1 := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{n1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
