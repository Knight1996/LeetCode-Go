//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。 
//
// 说明：本题中，我们将空字符串定义为有效的回文串。 
//
// 
//
// 示例 1: 
//
// 
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
// 
//
// 示例 2: 
//
// 
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 2 * 10⁵ 
// 字符串 s 由 ASCII 字符组成 
// 
// Related Topics 双指针 字符串 👍 513 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	//--------------------------------双指针---------------------------------
	// Time: O(n)
	// Space: O(1)
	s = strings.ToLower(s)
	left, right := 0, len(s) - 1
	for left < right {
		for left < right && !judge(s[left]) {
			left++
		}
		for left < right && !judge(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
	//----------------------------------------------------------------------
}

func judge(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

//leetcode submit region end(Prohibit modification and deletion)
