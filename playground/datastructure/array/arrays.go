package array

import "sort"

func SameArray(numbers1 []int, numbers2 []int) bool {
	if len(numbers1) != len(numbers2) {
		return false
	}

	for i := 0; i < len(numbers1); i++ {
		if numbers1[i] != numbers2[i] {
			return false
		}
	}

	return true
}

func ReverseArray(numbers []int) {
	start := 0
	end := len(numbers) - 1

	for start < end {
		numbers[start], numbers[end] = numbers[end], numbers[start]
		start++
		end--
	}
}

// https://leetcode.com/problems/sort-array-by-parity/
func SortArrayByParity(numbers []int) {
	var index int

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			numbers[index], numbers[i] = numbers[i], numbers[index]
			index++
		}
	}
}

// https://leetcode.com/problems/peak-index-in-a-mountain-array/
func PeakIndex(arr []int) int {
	max := 0
	peak := -1

	for i, num := range arr {
		if num > max {
			max = num
			peak = i
		}
	}

	return peak
}

// https://leetcode.com/problems/maximum-ascending-subarray-sum/
func MaxAscendingSum(nums []int) int {
	max := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}

		if currentSum > max {
			max = currentSum
		}
	}

	return max
}

// https://leetcode.com/problems/product-of-array-except-self/
func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	temp := 1
	for i := 1; i <= len(nums); i++ {
		result[i-1] = temp
		temp *= nums[i-1]
	}

	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= temp
		temp *= nums[i]
	}

	return result
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, i}
		} else {
			m[num] = i
		}
	}
	return nil
}

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}
