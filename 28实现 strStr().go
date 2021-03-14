// +ignore
package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issi"))
}
//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	neSize := len(needle)
	remainSize := len(haystack) - neSize
	for i := 0; i <= remainSize; i++ {
		if haystack[i:i+neSize] == needle {
			return i
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 19:02:19
//实现 strStr() 函数。 
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如
//果不存在，则返回 -1。 
//
// 示例 1: 
//
// 输入: haystack = "hello", needle = "ll"
//输出: 2
// 
//
// 示例 2: 
//
// 输入: haystack = "aaaaa", needle = "bba"
//输出: -1
// 
//
// 说明: 
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。 
// Related Topics 双指针 字符串 
// 👍 729 👎 0
