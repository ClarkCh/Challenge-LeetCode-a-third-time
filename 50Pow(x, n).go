// +ignore
package main

func main() {
	//fmt.Println(myPow(3, 5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	symbol := 1
	if n < 0 {
		n = -n
		symbol = -1
	} else if n == 0 {
		return 1
	}
	res := 1.0
	rec := x
	for n != 0 {
		if n&1 == 1 {
			res *= rec
		}
		rec *= rec
		n >>= 1
	}
	if symbol == -1 {
		res = 1 / res
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 17:18:53
//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
//
//
//
// 示例 1：
//
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//
//
// 示例 2：
//
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//
//
// 示例 3：
//
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
//
//
//
//
// 提示：
//
//
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// -104 <= xn <= 104
//
// Related Topics 数学 二分查找
// 👍 608 👎 0
