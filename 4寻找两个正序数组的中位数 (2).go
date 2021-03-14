// +ignore
package main

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cal(longNums, shortNums []int, index int) (float64, bool) {
	sizeMax := len(longNums) + len(shortNums)
	shortIndex := sizeMax - (sizeMax >> 1) - index
	longL := longNums[index-1]
	shortL := shortNums[shortIndex-1]
	longR := longNums[index]
	shortR := shortNums[shortIndex]
	if longL > shortR {
		return -1, false
	}
	if longR < shortL {
		return 1, false
	}
	if sizeMax & 1 == 1 {
		return float64(max(longL, shortL)), true
	}
	return (float64(max(longL, shortL)) + float64(min(longR, shortR))) / 2, true
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sizeMax, size1, size2, shortSize, longSize int
	var longNums, shortNums []int
	size1 = len(nums1)
	size2 = len(nums2)
	sizeMax = size1 + size2
	if size1 > size2 {
		longNums = nums1
		shortNums = nums2
		longSize = size1
		shortSize = size2
	}else {
		longNums = nums2
		shortNums = nums1
		longSize = size2
		shortSize = size1
	}
	// 先排除存在空数组的情况
	if shortSize < 1 {
		tmp := float64(longNums[longSize>>1])
		if sizeMax & 1 == 1 {

			return tmp
		}
		return (float64(longNums[longSize>>1-1]) + tmp) / 2
	}
	// 排除排除短数组的最左边和最右边
	// 在最左边
	tmp := longNums[sizeMax - (sizeMax>>1)-1]
	if tmp <= shortNums[0] {
		if size1 == size2 {
			return (float64(tmp) + float64(shortNums[0])) / 2
		}
		if sizeMax & 1 == 1 {
			return float64(tmp)
		}
		return (float64(tmp) + float64(min(shortNums[0], longNums[sizeMax - (sizeMax>>1)]))) / 2
	}
	// 在最右边
	tmp = longNums[longSize-(sizeMax>>1)]
	shortMax := shortNums[shortSize-1]
	if tmp >= shortNums[shortSize-1] {
		if size1 == size2 {
			return (float64(tmp) + float64(shortMax)) / 2
		}
		if sizeMax & 1 == 1 {
			return float64(max(shortMax, longNums[longSize-(sizeMax>>1)-1]))
		}
		return (float64(max(shortMax, longNums[longSize-(sizeMax>>1)-1])) + float64(tmp)) / 2
	}
	// 从长数组的[l, r]开始二分搜索
	l := longSize - (sizeMax >> 1) + 1
	r := sizeMax - (sizeMax>>1) - 1
	for l <= r {
		mid := (l + r) >> 1
		value, ok := cal(longNums, shortNums, mid)
		if ok {
			return value
		}
		if value > 0 {
			l = mid + 1
		}else if value < 0 {
			r = mid - 1
		}
	}
	return 0
}
//leetcode submon end(Prohibit modification and deletion)


// 2021-03-02 19:57:43
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
// 👍 3755 👎 0
