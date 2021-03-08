// +ignore
package main

func main() {
	nextPermutation([]int{1, 2, 3})
}

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int)  {
	size := len(nums)
	if size < 2 {
		return
	}
	var rec []int
	recI := size - 2
	for ; recI >= 0; recI-- {
		if nums[recI] < nums[recI+1] {
			rec = nums[recI:]
			break
		}
	}
	if rec == nil {
		tmpSize := size - 1
		for i := 0; i < (size >> 1); i++ {
			nums[i], nums[tmpSize-i] = nums[tmpSize-i], nums[i]
		}
		return
	}
	tmpSize := len(rec) - 1
	tmpNums := rec[1:]
	for i := 0; i < (tmpSize >> 1); i++ {
		tmpNums[i], tmpNums[tmpSize-1-i] = tmpNums[tmpSize-1-i], tmpNums[i]
	}
	for i := 0; i < tmpSize; i++ {
		if nums[recI] < tmpNums[i] {
			nums[recI], tmpNums[i] = tmpNums[i], nums[recI]
			return
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-08 16:29:22
//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。 
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。 
//
// 必须 原地 修改，只允许使用额外常数空间。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[1,3,2]
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,2,1]
//输出：[1,2,3]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,1,5]
//输出：[1,5,1]
// 
//
// 示例 4： 
//
// 
//输入：nums = [1]
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 100 
// 
// Related Topics 数组 
// 👍 987 👎 0


