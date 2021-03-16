// +ignore
package main

import (
	"strings"
)

func main() {
	isNumber("2e0")
}

//leetcode submit region begin(Prohibit modification and deletion)

func isInt(s string) bool {
	index := 0
	extend := false
	hasE := false
	size := len(s)
	for i, v := range s {

		if v == 101 || v == 69 {
			if hasE {
				return false
			}
			hasE = true
			index = i
			if i == 0 {
				return false
			}
			if s[i-1] < 48 || s[i-1] >= 58 {
				return false
			}
			if i+1 < size {
				if s[i+1] == 43 || s[i+1] == 45 {
					extend = true
					if i+2 < size && s[i+2] >= 48 && s[i+2] < 58 {

					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	tmp := index + 2
	if extend {
		tmp++
	}
	s2 := s[:index-1] + "2"
	if tmp < len(s) {
		s2 += s[tmp:]
	}
	for _, v := range s2 {
		if v < 48 || v >= 58 {
			return false
		}
	}
	return true
}

type listNode struct {
	next     map[byte]*listNode
	isNumber bool
}

func isNumber(s string) bool {
	size := len(s)
	if size < 1 {
		return false
	}
	if s[0] == 43 || s[0] == 45 {
		s = s[1:]
	}
	if strings.Count(s, ".") > 1 {
		return false
	}
	blank := false
	for i, v := range s {
		if v == 46 {
			if !isInt(s[:i]) && i != 0 {
				return false
			}
			if i == 0 {
				blank = true
			}
			s = s[i+1:]
			break
		}
	}
	if !isInt(s) && len(s) != 0 {
		return false
	}
	if blank && len(s) == 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-16 11:28:11
//有效数字（按顺序）可以分成以下几个部分：
//
//
// 一个 小数 或者 整数
// （可选）一个 'e' 或 'E' ，后面跟着一个 整数
//
//
// 小数（按顺序）可以分成以下几个部分：
//
//
// （可选）一个符号字符（'+' 或 '-'）
// 下述格式之一：
//
// 至少一位数字，后面跟着一个点 '.'
// 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
// 一个点 '.' ，后面跟着至少一位数字
//
//
//
//
// 整数（按顺序）可以分成以下几个部分：
//
//
// （可选）一个符号字符（'+' 或 '-'）
// 至少一位数字
//
//
// 部分有效数字列举如下：
//
//
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1",
// "53.5e93", "-123.456e789"]
//
//
// 部分无效数字列举如下：
//
//
// ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
//
//
// 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。
//
//
//
// 示例 1：
//
//
//输入：s = "0"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "e"
//输出：false
//
//
// 示例 3：
//
//
//输入：s = "."
//输出：false
//
//
// 示例 4：
//
//
//输入：s = ".1"
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。
//
// Related Topics 数学 字符串
// 👍 176 👎 0
