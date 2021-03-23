// +ignore
package main

import "fmt"

func main() {
	fmt.Println(countRangeSum([]int{2147483647, -2147483648, -1, 0}, -1, 0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func left2ISum(magic []int, i int) int {
	//fmt.Println(i)
	res := 0
	// 实现有问题
	reference := 1
	//for reference <= i {
	//	reference <<= 1
	//}
	//reference--
	//r2 := 0
	//fmt.Println("this:", i)
	//for reference ^ r2 > 0 {
	//	fmt.Println(i&(reference^r2))
	//	res += magic[i&(reference^r2)]
	//	if r2 > 0 {
	//		r2 <<= 1
	//	}else {
	//		r2 = 1
	//	}
	//reference >>= 1
	//}
	return res
}

func countRangeSum(nums []int, lower int, upper int) int {
	size := len(nums)
	magic := make([]int, size+1)
	res := 0
	for i := 1; i <= size; i++ {
		if nums[i-1] >= lower && nums[i-1] <= upper {
			res++
		}
		magic[i] += nums[i-1]
		tmp := i - (i & (i - 1))
		for j := 1; j < tmp; {
			magic[i] += magic[i-j]
			j <<= 1
		}
	}

	fmt.Println(magic)

	for i := 1; i <= size; i++ {
		for j := i + 1; j <= size; j++ {
			tmp := left2ISum(magic, j) - left2ISum(magic, i-1)
			if tmp >= lower && tmp <= upper {
				fmt.Println("j:", j, "i:", i, "tmp:", tmp)
				res++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-23 15:37:52
//给定一个整数数组 nums 。区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
//
// 请你以下标 i （0 <= i <= nums.length ）为起点，元素个数逐次递增，计算子数组内的元素和。
//
// 当元素和落在范围 [lower, upper] （包含 lower 和 upper）之内时，记录子数组当前最末元素下标 j ，记作 有效 区间和 S(i,
// j) 。
//
// 求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 有效 区间和的个数。
//
//
//
// 注意：
//最直观的算法复杂度是 O(n2) ，请在此基础上优化你的算法。
//
//
//
// 示例：
//
//
//输入：nums = [-2,5,-1], lower = -2, upper = 2,
//输出：3
//解释：
//下标 i = 0 时，子数组 [-2]、[-2,5]、[-2,5,-1]，对应元素和分别为 -2、3、2 ；其中 -2 和 2 落在范围 [lower =
//-2, upper = 2] 之间，因此记录有效区间和 S(0,0)，S(0,2) 。
//下标 i = 1 时，子数组 [5]、[5,-1] ，元素和 5、4 ；没有满足题意的有效区间和。
//下标 i = 2 时，子数组 [-1] ，元素和 -1 ；记录有效区间和 S(2,2) 。
//故，共有 3 个有效区间和。
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10^4
//
// Related Topics 排序 树状数组 线段树 二分查找 分治算法
// 👍 313 👎 0
