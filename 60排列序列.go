// +ignore
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(getPermutation(3, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)

func findNotUsed(dp []int, reference []bool, n, k int) byte {
	up := k / dp[n-1]
	if up*dp[n-1] != k {
		up++
	}
	cnt := 0 // 1
	for i, v := range reference {
		if !v {
			cnt++
			if cnt == up {
				reference[i] = true
				return byte(i + 49)
			}
		}
	}
	return 0
}

func helper(res *[]byte, dp []int, reference []bool, n, k int) {
	if k == 0 {
		up := n
		for i := len(reference) - 1; i >= 0; i-- {
			if !reference[i] {
				*res = append(*res, byte(i+49))
				reference[i] = true
				up--
				if up == 0 {
					return
				}
			}
		}
	}
	if n < 2 {
		for i, v := range reference {
			if !v {
				*res = append(*res, byte(i+49))
				return
			}
		}
		return
	}
	*res = append(*res, findNotUsed(dp, reference, n, k))
	helper(res, dp, reference, n-1, k%dp[n-1])
}

func getPermutation(n int, k int) string {
	dp := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	reference := make([]bool, n, n)
	res := make([]byte, 0, n)
	helper(&res, dp, reference, n, k)
	return *(*string)(unsafe.Pointer(&res))
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-16 14:40:51
//给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
//
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
//
//
// 给定 n 和 k，返回第 k 个排列。
//
//
//
// 示例 1：
//
//
//输入：n = 3, k = 3
//输出："213"
//
//
// 示例 2：
//
//
//输入：n = 4, k = 9
//输出："2314"
//
//
// 示例 3：
//
//
//输入：n = 3, k = 1
//输出："123"
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
// 1 <= k <= n!
//
// Related Topics 数学 回溯算法
// 👍 492 👎 0
