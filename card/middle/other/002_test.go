package other

import "testing"

func Test_evalRPN(t *testing.T) {

	input1 := []string{"2", "1", "+", "3", "*"}
	want1 := 9

	input2 := []string{"4", "13", "5", "/", "+"}
	want2 := 6

	input3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	want3 := 22

	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{input1}, want1},
		{"2", args{input2}, want2},
		{"3", args{input3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
