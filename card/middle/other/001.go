package other

import (
	"reflect"
)

// Sum of Two Integers
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/822/

func getSum(a int, b int) int {
	sum := 0
	carry := 0
	bits := uint(reflect.TypeOf(a).Bits())
	for i := uint(0); i < bits; i++ {
		bita := a & (1 << i)
		bitb := b & (1 << i)
		sum |= bita ^ bitb ^ carry
		carry = ((bita & bitb) | (bita & carry) | (bitb & carry)) << 1
	}

	return sum
}
