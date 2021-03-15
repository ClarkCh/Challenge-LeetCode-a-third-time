// +ignore
package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	i := 0
	last := len(nums) - 1
	for dp[last] == -1 {
		for j := 1; j <= nums[i]; j++ {
			if i+j > last {
				return dp[last]
			}
			if dp[i+j] == -1 {
				dp[i+j] = dp[i] + 1
			}
		}
		i++
	}
	return dp[last]
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 10:43:58
//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 示例:
//
// 输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 说明:
//
// 假设你总是可以到达数组的最后一个位置。
// Related Topics 贪心算法 数组
// 👍 863 👎 0
