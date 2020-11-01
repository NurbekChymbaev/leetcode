package _198_house_robber

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var memo = make([]int, len(nums)+1)

	memo[0] = 0
	memo[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		if memo[i] > memo[i-1]+nums[i] {
			memo[i+1] = memo[i]
		} else {
			memo[i+1] = memo[i-1] + nums[i]
		}
	}

	return memo[len(nums)]
}
