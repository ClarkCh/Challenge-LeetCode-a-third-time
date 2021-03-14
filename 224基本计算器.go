// +ignore
package main

func main() {
	//fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
//leetcode submit region begin(Prohibit modification and deletion)
var stringMap = make([]int, 58)
func sub(a, b int) int {
	return a - b
}
func add(a, b int) int {
	return a + b
}
func matchBracketFunc(s string, index int) (int, int) {
	// (2+3+(2+3)) -> 2+3+(2+3)
	size := len(s)
	res := index
	rec := 0  // rec是左括号的数目
	for index < size {
		switch s[index] {
		case 40:
			rec++
		case 41:
			rec--
		}
		if rec == 0 {
			return res, index
		}
		index++
	}
	return res, size - 1
}

func bracketFunc(s string) int {
	size := len(s)
	if size < 1 {
		return 0
	}

	var calFunc func(int, int) int = add
	res := 0
	rec := 0  // 用于储存未组合好的数字
	for i := 0; i < size; {
		switch s[i] {
		case 32:

		case 43:
			res = calFunc(res, rec)
			calFunc = add
			rec = 0
		case 45:
			res = calFunc(res, rec)
			calFunc = sub
			rec = 0
		case 40:
			//res = calFunc(res, rec)
			l, r := matchBracketFunc(s, i)  // [l, r] -> ()的索引
			res = calFunc(res, bracketFunc(s[l+1:r]))
			i = r  // 由于最后有个i++这里就不需要写r+1
			rec = 0
		default:
			// 目前预测是数字的时候
			rec = rec * 10 + stringMap[s[i]]
		}
		i++
	}
	return calFunc(res, rec)
}

func calculate(s string) int {
	for i := 0; i < 10; i++ {
		stringMap[i+48] = i
	}
	return bracketFunc(s)
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-10 22:52:05
//实现一个基本的计算器来计算一个简单的字符串表达式 s 的值。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "1 + 1"
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：s = " 2-1 + 2 "
//输出：3
// 
//
// 示例 3： 
//
// 
//输入：s = "(1+(4+5+2)-3)+(6+8)"
//输出：23
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 3 * 105 
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成 
// s 表示一个有效的表达式 
// 
// Related Topics 栈 数学 
// 👍 477 👎 0
