// +ignore
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("xabbcbbay"))
	fmt.Println(longestPalindrome("bb"))
}
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}
	b := make([]byte, 1, 2 * size - 1)
	b[0] = 35
	for _, v := range s {
		b = append(b, byte(v))
		b = append(b, 35)
	}
	s2 := string(b)
	dp := make([]int, 2 * size + 1)
	center := 1
	right := 2
	size2 := len(dp)
	dp[1] = 1
	//right]
	for i := 2; i < size2; i++ {
		if i <= right {
			dp[i] = dp[2*center-i]
			//if dp[i] + i < right {
			//
			//}
			if i + dp[i] >= right {
				dp[i] = right - i
				center = i
				//right++
				left := 2 * i - right
				for left > 0 && right + 1 < size2 && s2[left-1] == s2[right+1] {
					right++
					left--
					dp[i]++
				}

			}
		}else {
			center = i
			right = i
			left := i
			for left > 0 && right + 1 < size2 && s2[left-1] == s2[right+1] {
				right++
				left--
				dp[i]++
			}
		}
	}
	res := 0
	rec := 0
	for i, v := range dp {
		if v > res {
			res = v
			rec = i
		}
	}

	ret := make([]byte, 0)
	for i, v := range string(s2[rec-res:rec+res+1]) {
		if i & 1 == 1 {
			ret = append(ret, byte(v))
		}
	}
	return string(ret)
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-05 23:17:06
//给你一个字符串 s，找到 s 中最长的回文子串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 
//
// 示例 2： 
//
// 
//输入：s = "cbbd"
//输出："bb"
// 
//
// 示例 3： 
//
// 
//输入：s = "a"
//输出："a"
// 
//
// 示例 4： 
//
// 
//输入：s = "ac"
//输出："a"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 仅由数字和英文字母（大写和/或小写）组成 
// 
// Related Topics 字符串 动态规划 
// 👍 3287 👎 0
