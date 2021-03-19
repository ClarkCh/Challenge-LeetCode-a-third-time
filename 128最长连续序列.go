// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	ufParents := make(map[int]int)
	ufCnt := make(map[int]int)
	for _, v := range nums {
		ufParents[v] = v
		ufCnt[v] = 1
	}
	for _, v := range nums {
		if v2, ok := ufParents[v+1]; ok {
			if ufParents[v] > v2 {
				ufParents[v+1] = ufParents[v]
			} else {
				ufParents[v] = v2
			}
		}
	}
	res := 0
	for _, v := range nums {
		key := v
		for ufParents[key] != key {
			ufParents[key] = ufParents[ufParents[key]]
			key = ufParents[key]
		}
		tmp := key - v + 1
		if res < tmp {
			res = tmp
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-18 09:11:47
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//
//
// 进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 104
// -109 <= nums[i] <= 109
//
// Related Topics 并查集 数组
// 👍 715 👎 0
