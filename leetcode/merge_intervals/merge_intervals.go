package merge_intervals

import "sort"

// description
// https://leetcode.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	fixLeft := intervals[0][0]
	fixRight := intervals[0][1]
	for _, current := range intervals {
		if current[0] >= fixLeft && current[0] <= fixRight {
			if current[1] > fixRight {
				fixRight = current[1]
			}
		} else {
			merged = append(merged, []int{fixLeft, fixRight})
			fixLeft = current[0]
			fixRight = current[1]
		}
	}
	merged = append(merged, []int{fixLeft, fixRight})
	return merged
}
