// +ignore
package main

import (
	"fmt"
	"sort"
)

func main() {
	//fourSum([]int{-2, -1, 0, 0, 1, 2}, 0)
	fmt.Println(fourSum([]int{1,-2,-5,-4,-3,3,3,5}, -11))
}
//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	if size < 4 {
		return [][]int{}
	}
	res := make([][]int, 0)
	var ll, move, l, r int
	llLast := nums[ll] - 1

	tmpSum := 0
	for ll = 0; ll < size - 3; ll++ {
		if llLast == nums[ll] {
			continue
		}
		moveLast := nums[0] - 1
		for move = ll + 1; move < size - 2; move++ {
			if moveLast == nums[move] {
				continue
			}
			l = move + 1
			r = size - 1
			tmpSum = target - nums[ll] - nums[move]
			last := nums[0] - 1
			for l < r {
				if last == nums[l] {
					l++
					if last == nums[r] {
						break
					}
					continue
				}
				if last == nums[r] {
					r--
					if last == nums[l] {
						break
					}
					continue
				}
				if nums[l] + nums[r] > tmpSum {

					last = nums[r]
					r--
				}else if nums[l] + nums[r] < tmpSum {
					last = nums[l]
					l++

				}else {
					last = nums[r]
					res = append(res, []int{nums[ll], nums[move], nums[l], nums[r]})
					r--
				}
			}

			moveLast = nums[move]
		}
		llLast = nums[ll]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 00:42:31
//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。 
//
// 注意：答案中不可以包含重复的四元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [], target = 0
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 200 
// -109 <= nums[i] <= 109 
// -109 <= target <= 109 
// 
// Related Topics 数组 哈希表 双指针 
// 👍 762 👎 0
