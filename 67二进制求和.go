// +ignore
package main

import "unsafe"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	var shortString, longString string
	if len(a) > len(b) {
		shortString = b
		longString = a
	} else {
		shortString = a
		longString = b
	}
	res := make([]byte, 0, len(longString)+1)
	var rec byte = 0
	tmpLong := len(longString) - 1
	tmpShort := len(shortString) - 1
	for i := 0; i <= tmpShort; i++ {
		tmp := longString[tmpLong-i] + shortString[tmpShort-i] + rec - 96
		rec = tmp >> 1
		res = append(res, tmp&1+48)
	}
	for i := len(longString) - len(shortString) - 1; i >= 0; i-- {
		tmp := longString[i] + rec - 48
		rec = tmp >> 1
		res = append(res, tmp&1+48)
	}
	if rec == 1 {
		res = append(res, 49)
	}
	tmpSize := len(res) - 1
	for i := 0; i < len(res)>>1; i++ {
		res[i], res[tmpSize-i] = res[tmpSize-i], res[i]
	}
	return *(*string)(unsafe.Pointer(&res))
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-16 11:12:30
//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
// 输入: a = "11", b = "1"
//输出: "100"
//
// 示例 2:
//
// 输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//
// 提示：
//
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
//
// Related Topics 数学 字符串
// 👍 577 👎 0
