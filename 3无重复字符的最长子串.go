// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	l := 0
	res := 0
	cur := 0
	for i, v := range s {
		if value, ok := charMap[v]; ok {
			if value < l {
				cur++
			}else {
				res = max(res, cur)
				cur = i - value
				l = value
			}
		}else {
			cur++
		}

		charMap[v] = i
	}
	res = max(res, cur)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-01 15:53:38
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 
//
// 示例 2: 
//
// 
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 
//
// 示例 3: 
//
// 
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 示例 4: 
//
// 
//输入: s = ""
//输出: 0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 104 
// s 由英文字母、数字、符号和空格组成 
// 
// Related Topics 哈希表 双指针 字符串 Sliding Window 
// 👍 5037 👎 0


