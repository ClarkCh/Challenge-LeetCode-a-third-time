// +ignore
package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}
//leetcode submit region begin(Prohibit modification and deletion)
func helper(digits []int, reference []string, index, size int, rec []rune, res []string) []string {
	if index >= size {
		res = append(res, string(rec))
		return res
	}
	for _, v := range reference[digits[index]] {
		rec[index] = v
		res = helper(digits, reference, index+1, size, rec, res)
	}
	return res
}


func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	reference := []string{"", "", "abc", "def", "ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz"}
	intMap := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9}
	size := len(digits)
	rec := make([]rune, size, size)
	res := make([]string, 0)
	intDigits := make([]int, 0, size)
	for _, v := range digits {
		intDigits = append(intDigits, intMap[string(v)])
	}
	return helper(intDigits, reference, 0, size, rec, res)
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 00:20:27
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。 
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 
//
// 示例 2： 
//
// 
//输入：digits = ""
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：digits = "2"
//输出：["a","b","c"]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= digits.length <= 4 
// digits[i] 是范围 ['2', '9'] 的一个数字。 
// 
// Related Topics 深度优先搜索 递归 字符串 回溯算法 
// 👍 1164 👎 0
