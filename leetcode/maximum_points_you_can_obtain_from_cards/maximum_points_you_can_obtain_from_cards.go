package maximum_points_you_can_obtain_from_cards

// description https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	scoreLeft := make([]int, k+1)
	scoreRight := make([]int, k+1)

	for i := 1; i < k+1; i++ {
		scoreRight[i] = scoreRight[i-1] + cardPoints[n-i]
		scoreLeft[i] = scoreLeft[i-1] + cardPoints[i-1]
	}

	max := scoreRight[k]
	j := k
	for i := 0; i < k; i++ {
		if scoreRight[i]+scoreLeft[j] > max {
			max = scoreRight[i] + scoreLeft[j]
		}
		j--
	}
	return max
}
