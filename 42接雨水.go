// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	size := len(height)
	if size < 2 {
		return 0
	}
	l, r := 0, size-1
	lMax, rMax := 0, 0
	lMaxL, rMaxR := 0, 0
	has := 0
	sum := 0
	lUp, rUp := false, false
	for l < r {
		if !lUp {
			if height[l] > lMax {
				lMax = height[l]
				lMaxL = l
				lUp = !lUp
			} else {
				l++
				continue
			}
		}
		if !rUp {
			if height[r] > rMax {
				rMax = height[r]
				rMaxR = r
				rUp = !rUp
			} else {
				r--
				continue
			}
		}
		weidth := rMaxR - lMaxL + 1
		tmp := 0
		if rMax > lMax {
			tmp = lMax
			lUp = !lUp
		} else if lMax > rMax {
			tmp = rMax
			rUp = !rUp
		} else {
			tmp = rMax
			rUp = !rUp
			lUp = !lUp
		}
		sum += weidth * (tmp - has)
		has = tmp

	}

	for _, v := range height {
		if v > has {
			sum -= has
		} else {
			sum -= v
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-14 12:00:08
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 0 <= n <= 3 * 104
// 0 <= height[i] <= 105
//
// Related Topics 栈 数组 双指针 动态规划
// 👍 2132 👎 0
