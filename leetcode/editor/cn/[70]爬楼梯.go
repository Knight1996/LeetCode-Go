//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//
// 示例 2：
//
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
//
//
//
// 提示：
//
//
// 1 <= n <= 45
//
// Related Topics 记忆化搜索 数学 动态规划 👍 2274 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	//----------------------------动态规划--------------------------------
	/*// Time: O(n)
	// Space: O(n)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int , n + 1)
	dp[1], dp[2] = 1, 2

	for i := 3 ; i <= n ; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]*/
	//-------------------------------------------------------------------

	//-------------------------动态规划(空间优化)---------------------------
	// Time: O(n)
	// Space: O(1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 3 ; i <= n ; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
	//-------------------------------------------------------------------
}

func min(a , b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
