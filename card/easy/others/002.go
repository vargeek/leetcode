package others

// Hamming Distance
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/762/
// https://leetcode.com/problems/hamming-distance/description/

func hammingDistance(x int, y int) int {
	xor := x ^ y
	distance := 0
	for xor != 0 {
		distance += xor & 0x01
		xor >>= 1
	}
	return distance
}
