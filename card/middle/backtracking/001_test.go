package backtracking

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	digits1 := "23"
	want1 := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{digits1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
