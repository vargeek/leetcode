package hard

// 4Sum II
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/829/
// https://leetcode.com/submissions/detail/173751898/

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sum1 := make(map[int]int, len(A)*len(B))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sum1[-A[i]-B[j]]++
		}
	}
	count := 0
	for k := 0; k < len(C); k++ {
		for l := 0; l < len(D); l++ {
			count += sum1[C[k]+D[l]]
		}
	}
	return count
}
