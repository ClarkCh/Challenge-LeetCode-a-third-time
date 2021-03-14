// +ignore
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(multiply("3", "2"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var longNumString, shortNumString string
	//longSize, shortSize := len(longNumString), len(shortNumString)
	if len(num1) > len(num2) {
		longNumString = num1
		shortNumString = num2
	} else {
		longNumString = num2
		shortNumString = num1
	}
	longSize, shortSize := len(longNumString), len(shortNumString)
	longNum := make([]int, len(longNumString))
	shortNum := make([]int, len(shortNumString))
	for i, v := range longNumString {
		longNum[i] = int(v - 48)
	}
	for i, v := range shortNumString {
		shortNum[i] = int(v - 48)
	}
	sum := make([]int, longSize+shortSize)
	tmpShortSize := shortSize - 1
	tmpLongSize := longSize - 1
	tmpSumSize := len(sum) - 1
	for i := 0; i < shortSize; i++ {
		for j := 0; j < longSize; j++ {
			sum[tmpSumSize-i-j] += shortNum[tmpShortSize-i] * longNum[tmpLongSize-j]
		}
	}
	tmp := 0
	for i := tmpSumSize; i >= 0; i-- {
		tmpSum := sum[i] + tmp
		tmp = tmpSum / 10
		sum[i] = tmpSum % 10
	}

	for i, v := range sum {
		if v != 0 {
			sum = sum[i:]
			break
		}
	}

	res := make([]byte, len(sum))
	for i, v := range sum {
		res[i] = byte(v) + 48
	}
	return *(*string)(unsafe.Pointer(&res))
	//return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-14 13:09:46
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 示例 1:
//
// 输入: num1 = "2", num2 = "3"
//输出: "6"
//
// 示例 2:
//
// 输入: num1 = "123", num2 = "456"
//输出: "56088"
//
// 说明：
//
//
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
// Related Topics 数学 字符串
// 👍 579 👎 0
