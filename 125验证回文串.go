// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	dictMap := make(map[byte]byte)
	var key byte
	for key = 48; key < 58; key++ {
		dictMap[key] = key
	}
	for key = 65; key < 91; key++ {
		dictMap[key] = key
	}
	for key = 97; key < 123; key++ {
		dictMap[key] = key - 32
	}
	s2b := make([]byte, 0, len(s))
	for i := range s {
		if v, ok := dictMap[s[i]]; ok {

			s2b = append(s2b, v)
		}
	}
	size := len(s2b)
	up := size >> 1
	tmp := size - 1
	for i := 0; i < up; i++ {
		if s2b[i] != s2b[tmp-i] {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-01 15:53:08
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。 
//
// 说明：本题中，我们将空字符串定义为有效的回文串。 
//
// 示例 1: 
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "race a car"
//输出: false
// 
// Related Topics 双指针 字符串 
// 👍 340 👎 0


