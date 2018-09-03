package arrayandstring

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	want1 := [][]string{
		[]string{"eat", "tea", "ate"},
		[]string{"tan", "nat"},
		[]string{"bat"},
	}
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"0", args{strs1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}

}
