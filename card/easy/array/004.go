package easy

// 存在重复
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/24/
// https://leetcode-cn.com/problems/contains-duplicate/description/

func quickSort(nums []int) {
	var _loop func([]int, int, int)
	_loop = func(nums []int, left int, right int) {
		if right <= left {
			return
		}

		i := left
		j := right
		m := (left + right) / 2
		val := nums[m]
		for i <= j {
			for nums[i] < val && i <= j {
				i++
			}
			for nums[j] > val && i <= j {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		_loop(nums, left, j)
		_loop(nums, i, right)
	}

	_loop(nums, 0, len(nums)-1)
}

func containsDuplicate1(nums []int) bool {
	// sort.Ints(nums)
	quickSort(nums)
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	var _loop func([]int, int, int) bool
	_loop = func(nums []int, left int, right int) bool {
		if right <= left {
			return false
		}

		i := left
		j := right
		val := nums[(left+right)/2]
		for i <= j {
			for nums[i] < val && i <= j {
				i++
			}
			for nums[j] > val && i <= j {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		return _loop(nums, left, j) || _loop(nums, i, right) || (j >= left && i <= right && nums[i] == nums[j])
	}
	return _loop(nums, 0, len(nums)-1)
}
