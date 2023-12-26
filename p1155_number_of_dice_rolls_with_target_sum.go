package main

func dp(n int, k int, target int, memo map[[2]int]int) int {
	res := 0
	key := [2]int{n, target}
	if val, ok := memo[key]; ok {
		return val
	}
	if n == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	if target <= 0 || target > n*k {
		return 0
	}

	for i := 1; i <= k; i++ {
		res += dp(n-1, k, target-i, memo)
	}
	res = res % 1000000007
	memo[key] = res

	return res

}

func numRollsToTarget(n int, k int, target int) int {
	memo := map[[2]int]int{}
	return dp(n, k, target, memo)
}
