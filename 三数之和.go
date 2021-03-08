// +ignore
package main

import (
	"car/leetcode/leetcode/editor/cn/utils"
	"sort"
)

func main() {
	threeSum(utils.Trans2Array("[-1,0,1,2,-1,-4]"))
}


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	res := make([][]int, 0)
	last := nums[0] - 1
	for mid := 0; mid < len(nums) - 2; mid++ {
		if last == nums[mid] {
			continue
		}
		l = mid + 1
		tmp := - nums[mid]
		leftHalfTmp := tmp / 2
		rightHalfTmp := 0
		if tmp % 2 == 0 {
			rightHalfTmp = leftHalfTmp
		}else {
			rightHalfTmp = leftHalfTmp + 1
		}

		for l < r {
			//if nums[l] * 2 > tmp || nums[r] * 2 < tmp {
			//	goto label
			//}
			if nums[l] > leftHalfTmp || nums[r] < rightHalfTmp {
				goto label
			}
			if nums[l] + nums[r] > tmp {
				r--
			}else if nums[l] + nums[r] < tmp {
				l++
			}else {
				res = append(res, []int{nums[mid], nums[l], nums[r]})
			}
		}
		label:
		last = nums[mid]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-01 14:05:27
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -105 <= nums[i] <= 105 
// 
// Related Topics 数组 双指针 
// 👍 3019 👎 0


