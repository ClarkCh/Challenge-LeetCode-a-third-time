// +ignore
package main

import "fmt"

func main() {
	//a := []int{1, 2, 3, 7, 8, 8}
	//b := []int{2, 4, 9, 12}
	a := []int{1, 3}
	b := []int{2}
	fmt.Println(findMedianSortedArrays(a, b))
}

//leetcode submit region begin(Prohibit modification and deletion)
func dropLeft(nums1, nums2 []int, K int) int {
	size1 := len(nums1)
	size2 := len(nums2)
	var longNums, shortNums []int
	var shortSize int
	if size1 < size2 {
		longNums = nums2
		shortNums = nums1
		shortSize = size1
	}else {
		longNums = nums1
		shortNums = nums2
		shortSize = size2
	}
	if shortSize < 1 {
		return longNums[K-1]
	}
	if K == 1 {
		if shortNums[0] > longNums[0] {
			return longNums[0]
		}
		return shortNums[0]
	}
	sub := K >> 1
	if shortSize < sub {
		sub = shortSize
	}
	if longNums[sub-1] < shortNums[sub-1] {
		return dropLeft(longNums[sub:], shortNums, K-sub)
	}
	return dropLeft(longNums, shortNums[sub:], K-sub)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sizeMax := len(nums1) + len(nums2)
	if sizeMax & 1 == 1 {
		return float64(dropLeft(nums1, nums2, sizeMax-(sizeMax>>1)))
	}
	return (float64(dropLeft(nums1, nums2, (sizeMax>>1)+1)) + float64(dropLeft(nums1, nums2, sizeMax>>1))) / 2
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-02 14:24:32
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 
//
// 示例 3： 
//
// 
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
// 
//
// 示例 4： 
//
// 
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
// 
//
// 示例 5： 
//
// 
//输入：nums1 = [2], nums2 = []
//输出：2.00000
// 
//
// 
//
// 提示： 
//
// 
// nums1.length == m 
// nums2.length == n 
// 0 <= m <= 1000 
// 0 <= n <= 1000 
// 1 <= m + n <= 2000 
// -106 <= nums1[i], nums2[i] <= 106 
// 
//
// 
//
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？ 
// Related Topics 数组 二分查找 分治算法 
// 👍 3754 👎 0


