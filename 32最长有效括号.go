// +ignore
package main

func main() {
	//longestValidParentheses("(()")
	//longestValidParentheses("()")
	longestValidParentheses("()(()")
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	//size := len(s)
	res := 0
	rec := 0
	left := 0
	for _, v := range s {
		if v == 40 {
			left++
			rec++
		}else {
			if left > 0 {
				rec++
				left--
			}else {
				if rec > res {
					res = rec
					rec = 0
					left = 0
				}
			}
		}
	}
	if left > 0 {
		// 这种情况下会需要剔除多出来的左括号左边的成对括号
	}
	//if left >= 0 {
		if rec - left > res {
			res = rec - left
		}
	//}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-08 16:50:14
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。 
//
// 
//
// 
// 
// 示例 1： 
//
// 
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
// 
//
// 示例 2： 
//
// 
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
// 
//
// 示例 3： 
//
// 
//输入：s = ""
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 3 * 104 
// s[i] 为 '(' 或 ')' 
// 
// 
// 
// Related Topics 字符串 动态规划 
// 👍 1199 👎 0


