package hard

import "testing"

func Test_fourSumCount(t *testing.T) {
	a1 := []int{1, 2}
	b1 := []int{-2, -1}
	c1 := []int{-1, 2}
	d1 := []int{0, 2}
	want1 := 2
	a2 := []int{0}
	b2 := []int{0}
	c2 := []int{0}
	d2 := []int{0}
	want2 := 1
	type args struct {
		A []int
		B []int
		C []int
		D []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1, b1, c1, d1}, want1},
		{"2", args{a2, b2, c2, d2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
