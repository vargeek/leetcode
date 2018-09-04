package hard

// Container With Most Water
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/830/
// https://leetcode.com/problems/container-with-most-water/

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
func maxArea1(height []int) int {
	area := 0
	for i, h := range height {
		for j := 0; j < i; j++ {
			if height[j] >= h {
				area = max((i-j)*h, area)
				break
			}
		}
		for j := len(height) - 1; j > i; j-- {
			if height[j] >= h {
				area = max((j-i)*h, area)
				break
			}
		}
	}
	return area
}

func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		area = max(area, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
