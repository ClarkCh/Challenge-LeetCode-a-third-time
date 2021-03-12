// +ignore
package main

//func main() {
//	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
//}

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {

	size := len(nums)
	if size < 1 {
		return 1
	}
	//hasOut := false
	//for _, v := range nums {
	//	if v == size {
	//		hasOut = true
	//		break
	//	}
	//}  // 和下一行二选一
	nums = append(nums, -1) // 把[2, 1, 3]中3的位置预留一下
	size = len(nums)
	needBeFill := 0
	i, tmp := -1, 0
	for {
		i++
		if i >= size {
			break
		}
		needBeFill = nums[i]
		if needBeFill == i {
			continue
		}
		nums[i] = -1
	label:
		if needBeFill < 0 || needBeFill >= size || nums[needBeFill] == needBeFill {
			continue
		}
		tmp = nums[needBeFill]
		nums[needBeFill] = needBeFill
		needBeFill = tmp
		goto label
	}
	for i = 1; i < size; i++ {
		if nums[i] != i {
			return i
		}
	}
	//if hasOut {
	//	return size + 1
	//}
	return len(nums)
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-12 18:28:02
//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
//
//
// 进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 300
// -231 <= nums[i] <= 231 - 1
//
// Related Topics 数组
// 👍 993 👎 0
