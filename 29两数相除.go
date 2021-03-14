// +ignore
package main

import "fmt"

func main() {
	fmt.Println(1 << 31)
	//fmt.Println(divide(10, 3))
	fmt.Println(divide(10, 1))

}
//leetcode submit region begin(Prohibit modification and deletion)
func helper(dividend, divesor int) int {
	if dividend >= 0 && dividend < divesor {
		return 0
	}
	remainder := 1
	curDivesor := divesor
	for dividend >= curDivesor {
		curDivesor <<= 1
		remainder <<= 1
	}
	remainder >>= 1
	return remainder + helper(dividend - (curDivesor >> 1), divesor)
}


func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	symbol := 1
	if divisor < 0 {
		symbol = -symbol
		divisor = -divisor
	}
	if dividend < 0 {
		symbol = -symbol
		dividend = -dividend
	}
	res := symbol * helper(dividend, divisor)
	if res >= 2147483647 || res < -2147483648 {
		return 2147483647
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 22:07:02
//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。 
//
// 返回被除数 dividend 除以除数 divisor 得到的商。 
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2 
//
// 
//
// 示例 1: 
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3 
//
// 示例 2: 
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2 
//
// 
//
// 提示： 
//
// 
// 被除数和除数均为 32 位有符号整数。 
// 除数不为 0。 
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。 
// 
// Related Topics 数学 二分查找 
// 👍 517 👎 0
