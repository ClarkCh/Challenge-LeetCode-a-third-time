// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	size := len(digits)
	res := 1
	for i := size - 1; i >= 0; i-- {
		if digits[i]+res == 10 {
			digits[i] = 0
			res = 1
		} else {
			digits[i]++
			res = 0
			break
		}
	}
	if res == 1 {
		res2 := make([]int, 1+len(digits))
		res2[0] = 1
		copy(res2[1:], digits)
		return res2
	}
	return digits
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-16 10:55:45
//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//
//
// 示例 1：
//
//
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
//
//
// 示例 2：
//
//
//输入：digits = [4,3,2,1]
//输出：[4,3,2,2]
//解释：输入数组表示数字 4321。
//
//
// 示例 3：
//
//
//输入：digits = [0]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9
//
// Related Topics 数组
// 👍 643 👎 0
