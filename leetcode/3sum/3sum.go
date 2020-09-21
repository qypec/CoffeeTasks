package sum

import (
	"sort"
	"strconv"
)

// description
// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)

	unique := make(map[string]bool)
	result := make([][]int, 0)
	sum := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			sum = nums[i] + nums[j]

			index := sort.SearchInts(nums[j+1:], -sum)
			if index+j+1 < len(nums) && nums[index+j+1] == -sum {
				resStr := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[index+j+1])
				if _, ok := unique[resStr]; !ok {
					result = append(result, []int{
						nums[i], nums[j], nums[index+j+1],
					})
					unique[resStr] = true
				}
			}
		}
	}
	return result
}
