// +ignore
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

//leetcode submit region begin(Prohibit modification and deletion)
func dfs(res *[][]int, candidates, rec []int, index, target int) {
	if index >= len(candidates) {
		return
	}
	//if target == 0 {
	//	*res = append(*res, rec)
	//	return
	//}
	rec = append(rec, 0)
	recIndex := len(rec) - 1
	for i := index; i < len(candidates); i++ {
		rec[recIndex] = candidates[i]
		if target == candidates[i] {
			tmp := make([]int, len(rec), len(rec))
			copy(tmp, rec)
			*res = append(*res, tmp)
			return
		} else if target < candidates[i] {
			return
		}
		dfs(res, candidates, rec, i, target-candidates[i])
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	dfs(&res, candidates, make([]int, 0), 0, target)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-12 14:12:15
//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。
//
// 说明：
//
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
//
//
// 示例 1：
//
// 输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
//
//
// 示例 2：
//
// 输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都是独一无二的。
// 1 <= target <= 500
//
// Related Topics 数组 回溯算法
// 👍 1211 👎 0
