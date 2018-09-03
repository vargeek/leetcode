package mmath

import (
	"fmt"
)

// Fraction to Recurring Decimal
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/821/

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	bytes := []byte{}
	if (numerator > 0) != (denominator > 0) {
		bytes = append(bytes, '-')
	}

	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}

	bytes = append(bytes, []byte(fmt.Sprintf("%d", numerator/denominator))...)
	numerator = numerator % denominator
	if numerator == 0 {
		return string(bytes)
	}

	bytes = append(bytes, '.')
	indices := map[int]int{}
	for numerator != 0 {
		if index, ok := indices[numerator]; ok {
			bytes = append(bytes, '(', ')')
			for i := len(bytes) - 2; i > index; i-- {
				bytes[i] = bytes[i-1]
			}
			bytes[index] = '('
			break
		} else {
			indices[numerator] = len(bytes)
			numerator *= 10
			bytes = append(bytes, []byte(fmt.Sprintf("%d", numerator/denominator))...)
			numerator %= denominator
		}
	}
	return string(bytes)
}
