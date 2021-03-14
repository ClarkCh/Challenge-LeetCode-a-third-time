// +ignore
package main

import "fmt"

func main() {
	fmt.Println("()"[0])
	fmt.Println("()"[1])
}
//leetcode submit region begin(Prohibit modification and deletion)

func helper(rec []byte, res []string, index, usedLeft, remainUseLeft, n, size int) []string {
	if index >= size {
		res = append(res, string(rec))
		return res
	}
	if usedLeft < n {
		rec[index] = 40
		res = helper(rec, res, index+1, usedLeft+1, remainUseLeft+1, n, size)
	}
	if remainUseLeft > 0 {
		rec[index] = 41
		res = helper(rec, res, index+1, usedLeft, remainUseLeft-1, n, size)
	}
	return res
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	rec := make([]byte, 2 * n)
	return helper(rec, res, 0, 0, 0, n, 2*n)
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 02:32:15
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：["()"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 8 
// 
// Related Topics 字符串 回溯算法 
// 👍 1596 👎 0
