//给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。 
//
// 你从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。 
//
// 如果路径在任何位置上与自身相交，也就是走到之前已经走过的位置，请返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：path = "NES"
//输出：false 
//解释：该路径没有在任何位置相交。 
//
// 示例 2： 
//
// 
//
// 
//输入：path = "NESWW"
//输出：true
//解释：该路径经过原点两次。 
//
// 
//
// 提示： 
//
// 
// 1 <= path.length <= 10⁴ 
// path[i] 为 'N'、'S'、'E' 或 'W' 
// 
// Related Topics 哈希表 字符串 👍 34 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPathCrossing(path string) bool {
	//------------------------------hashSet----------------------------------
	// Time: O(n)
	// Space: O(n)
	x, y := 0, 0
	hashSet := map[[2]int]bool{[2]int{0,0} : true}
	for i := 0 ; i < len(path) ; i++ {
		switch string(path[i]) {
			case "N":
				y++
			case "S":
			    y--
			case "E":
			    x++
			case "W":
				x--
		}
		if hashSet[[2]int{x , y}] == true {
			return true
		}
		hashSet[[2]int{x , y}] = true
	}
	return false
	//-------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
