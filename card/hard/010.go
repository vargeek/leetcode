package hard

// Basic Calculator II
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/836/
// https://leetcode.com/problems/basic-calculator-ii/

func calculate(s string) int {
	ops := map[rune]func(int, int) int{
		'+': func(op1, op2 int) int {
			return op1 + op2
		},
		'-': func(op1, op2 int) int {
			return op1 - op2
		},
		'*': func(op1, op2 int) int {
			return op1 * op2
		},
		'/': func(op1, op2 int) int {
			return op1 / op2
		},
	}

	priorities := map[rune]int{
		'=': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	opStack := make([]rune, 0)
	nums := make([]int, 0)
	num := 0

	pushOp := func(op rune) {
		priority := priorities[op]
		for len(opStack) > 0 {
			lastOp := opStack[len(opStack)-1]
			lastPri := priorities[lastOp]
			if lastPri < priority {
				break
			}
			l := len(nums)
			if fn, ok := ops[lastOp]; ok {
				nums[l-2] = fn(nums[l-2], nums[l-1])
				nums = nums[:len(nums)-1]
				opStack = opStack[:len(opStack)-1]
			} else {
				panic("Opps")
			}
		}
		opStack = append(opStack, op)
	}
	for _, ch := range s {
		switch ch {
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			nums = append(nums, num)
			num = 0
			pushOp(ch)
		case ' ':
		default:
			num = num*10 + int(ch-'0')
		}
	}
	nums = append(nums, num)
	pushOp('=')

	return nums[0]
}
