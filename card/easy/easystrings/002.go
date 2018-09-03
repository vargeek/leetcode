package easystrings

// 颠倒整数
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/33/
// https://leetcode-cn.com/problems/reverse-integer/

func reverse1(x int) int {
	val := int32(x)
	bits := []int32{}
	for val != 0 {
		bit := val % 10
		val /= 10
		bits = append(bits, bit)
	}

	val = int32(x)

	ans := int32(0)
	factor := int32(1)
	for i := len(bits) - 1; i >= 0; i-- {
		tmp := factor * bits[i]
		if tmp/factor != bits[i] {
			ans = 0
			break
		}
		ans += tmp
		if (ans ^ val) < 0 {
			ans = 0
			break
		}
		factor *= 10
	}
	return int(ans)
}

func reverse(x int) int {
	val := int32(x)
	ans := int32(0)
	sign := val

	for val != 0 {
		bit := val % 10
		tmp := ans * 10
		if tmp/10 != ans {
			return 0
		}
		ans = tmp + bit

		if ans != 0 && ans^sign < 0 {
			return 0
		}
		val /= 10
	}

	return int(ans)
}
