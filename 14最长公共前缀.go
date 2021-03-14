// +ignore
package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) < 2 {
		return strs[0]
	}
	minValue := len(strs[0])
	size := 0
	for _, v := range strs {
		size = len(v)
		if size < minValue {
			minValue = size
		}
	}
	up := 0
	reference := strs[0]
	for i := 0; i < minValue; i++ {
		for _, v := range strs[1:] {
			if reference[i] != v[i] {
				return reference[:up]
			}
		}
		up++
	}
	return reference[:up]
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-02 20:48:56
//编写一个函数来查找字符串数组中的最长公共前缀。 
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 
//
// 示例 2： 
//
// 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
// 
//
// 提示： 
//
// 
// 0 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 
// Related Topics 字符串 
// 👍 1482 👎 0
