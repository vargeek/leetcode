package other

import (
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 2}, 3},
		{"2", args{-2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}

	// foo := -1
	// size := reflect.TypeOf(foo).Size()
	// sign := (1 << (size*8 - 1))

	// fmt.Printf("-1>>1: %x\n", foo>>1)
	// fmt.Printf("%x, %x", ^foo, sign)
	// fmt.Println("\n________分割线——————————————————")
}
