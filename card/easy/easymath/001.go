package easymath

import "strconv"

// Fizz Buzz
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/743/
// https://leetcode.com/problems/fizz-buzz/description/

func fizzBuzz(n int) []string {
	texts := []string{}
	for i := 1; i <= n; i++ {
		fizz, buzz := i%3, i%5

		if fizz == 0 {
			if buzz == 0 {
				texts = append(texts, "FizzBuzz")
			} else {
				texts = append(texts, "Fizz")
			}
		} else if buzz == 0 {
			texts = append(texts, "Buzz")
		} else {
			texts = append(texts, strconv.Itoa(i))
		}
	}

	return texts
}
