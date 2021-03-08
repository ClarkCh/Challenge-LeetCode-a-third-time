// +ignore
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1,0,1,2,-1,-4}
	threeSum(a)
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	size := len(nums)
	if size < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	fmt.Println(nums)
	l, r, move := 1, size-1, 0
	lastMove := nums[move] - 1
	target := 0
	res := make([][]int, 0)
	for ; move < size - 2; move++ {
		if lastMove == nums[move] {
			continue
		}
		if nums[move] > 0 {
			break
		}
		l = move + 1
		r = size - 1
		target = - nums[move]
		last := nums[l] - 1
		for l < r {
			if nums[l] == last {
				l++
				continue
			}else if nums[r] == last {
				r--
				continue
			}
			sum := nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[l], nums[move], nums[r]})
				last = nums[r]
				r--
			}else if sum > target {
				last = nums[r]
				r--
			}else {
				last = nums[l]
				l++
			}
		}
		lastMove = nums[move]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-03 10:24:29
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
// 👍 3029 👎 0


