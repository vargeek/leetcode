package sortingandsearching

// Merge Sorted Array
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/587/
// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	total := m + n
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	idx1 := n
	idx2 := 0
	idx := 0

	for idx1 < total && idx2 < n {
		if nums1[idx1] < nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1++
		} else {
			nums1[idx] = nums2[idx2]
			idx2++
		}
		idx++
	}
	for idx1 < total {
		nums1[idx] = nums1[idx1]
		idx1++
		idx++
	}
	for idx2 < n {
		nums1[idx] = nums2[idx2]
		idx2++
		idx++
	}

	//     0 1 2 3
	// 0 1 2 3 4 5
	// 6 - 3
}
