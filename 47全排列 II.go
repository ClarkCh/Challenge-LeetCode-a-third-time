// +ignore
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func helper(nums, rec []int, reference []bool, res *[][]int) {
	size := len(nums)
	if len(rec) == size {
		tmp := make([]int, size, size)
		copy(tmp, rec)
		*res = append(*res, tmp)
	}
	for i := 0; i < size; i++ {
		if !reference[i] {
			if i > 0 && nums[i] == nums[i-1] && reference[i-1] {
				continue
			}

			rec = append(rec, nums[i])
			reference[i] = true
			helper(nums, rec, reference, res)
			reference[i] = false
			rec = rec[:len(rec)-1]
		}
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	rec := make([]int, 0, size)
	res := make([][]int, 0)
	reference := make([]bool, size)
	helper(nums, rec, reference, &res)
	//for i := 0; i < size; i++ {
	//	if i > 0 && nums[i] == nums[i-1] {
	//		continue
	//	}
	//	reference[i] = true
	//	rec = append(rec, nums[i])
	//	helper(nums, rec, reference, &res)
	//	reference[i] = false
	//	rec = rec[:0]
	//}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 11:16:25
//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 回溯算法
// 👍 629 👎 0
