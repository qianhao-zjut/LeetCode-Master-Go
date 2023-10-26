package main
//746 最小花费怕楼梯
func minCostClimbingStairs(cost []int) int {

	l := len(cost)
	var res int
	dp := make([]int, l, l)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < l; i++ {
		dp[i] = cost[i] + min(dp[i-1],dp[i-2])
	}
	
	res = min(dp[l-1],dp[l-2])
	
	return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}