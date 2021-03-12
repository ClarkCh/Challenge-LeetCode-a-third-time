// +ignore
package main

import "sort"

//func main() {
//	combinationSum2([]int{3,1,3,5,1,1}, 8)
//}
//leetcode submit region begin(Prohibit modification and deletion)
func dfs(res *[][]int, candidates, rec []int, index, target int) {
	if index >= len(candidates) {
		return
	}
	rec = append(rec, 0)
	recIndex := len(rec) - 1
	last := candidates[index] - 1
	for i := index; i < len(candidates); i++ {
		if last == candidates[i] {
			continue
		}
		rec[recIndex] = candidates[i]
		if target == candidates[i] {
			tmp := make([]int, len(rec), len(rec))
			copy(tmp, rec)
			*res = append(*res, tmp)
			return
		} else if target < candidates[i] {
			return
		}
		dfs(res, candidates, rec, i+1, target-candidates[i])
		last = candidates[i]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	rec := make([]int, 0)
	dfs(&res, candidates, rec, 0, target)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-12 18:12:51
//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。
//
// 说明：
//
//
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
//
//
// 示例 1:
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//
//
// 示例 2:
//
// 输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//]
// Related Topics 数组 回溯算法
// 👍 522 👎 0
