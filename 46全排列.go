// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func helper(nums, rec []int, reference []bool, res *[][]int) {
	size := len(nums)
	if len(rec) == size {
		tmp := make([]int, size, size)
		copy(tmp, rec)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < size; i++ {
		if !reference[i] {

			rec = append(rec, nums[i])
			reference[i] = true
			helper(nums, rec, reference, res)
			rec = rec[:len(rec)-1]
			reference[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	reference := make([]bool, len(nums))
	helper(nums, make([]int, 0, len(nums)), reference, &res)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 10:59:47
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法
// 👍 1203 👎 0
