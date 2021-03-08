// +ignore
package main

import "sort"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	size := len(nums)
	var l, move, r, reference int
	lastMove := nums[0] - 1
	res := nums[0] + nums[1] + nums[2]
	if res > target {
		reference = res - target
	}else {
		reference = target - res
	}
	for move = 0; move < size - 2; move++ {
		if lastMove == nums[move] {
			continue
		}
		l = move + 1
		r = size - 1

		tmpTarget := target - nums[move]
		for l < r{
			tmp := nums[l] + nums[r]
			if tmpTarget < tmp {
				if tmp - tmpTarget < reference {
					reference = tmp - tmpTarget
					res = tmp + nums[move]
				}
				r--
			}else if tmpTarget > tmp {
				if tmpTarget - tmp < reference {
					reference = tmpTarget - tmp
					res = tmp + nums[move]
				}
				l++
			}else {
				return target
			}


		}
		lastMove = nums[move]

	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-03 18:30:45
//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。 
//
// 
//
// 示例： 
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 10^3 
// -10^3 <= nums[i] <= 10^3 
// -10^4 <= target <= 10^4 
// 
// Related Topics 数组 双指针 
// 👍 702 👎 0


