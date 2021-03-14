// +ignore
package main

import "fmt"

func main() {
	fmt.Println("[](){}"[0])
	fmt.Println("[](){}"[1])
	fmt.Println("[](){}"[2])
	fmt.Println("[](){}"[3])
	fmt.Println("[](){}"[4])
	fmt.Println("[](){}"[5])
}
//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stringMap := make([]rune, 256, 256)
	stringMap[93] = 91
	stringMap[41] = 40
	stringMap[125] = 123
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if v == 91 || v == 40 || v == 123 {
			stack = append(stack, v)
		}else {
			if len(stack) < 1{
				return false
			}
			tmp := len(stack)-1
			if stack[tmp] == stringMap[v] {
				stack = stack[:tmp]
			}else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 01:26:04
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "()[]{}"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：s = "(]"
//输出：false
// 
//
// 示例 4： 
//
// 
//输入：s = "([)]"
//输出：false
// 
//
// 示例 5： 
//
// 
//输入：s = "{[]}"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 104 
// s 仅由括号 '()[]{}' 组成 
// 
// Related Topics 栈 字符串 
// 👍 2211 👎 0
