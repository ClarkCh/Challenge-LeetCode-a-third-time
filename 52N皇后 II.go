// +ignore
package main

func main() {
	totalNQueens(9)
}

//leetcode submit region begin(Prohibit modification and deletion)
func dfs(res *int, cols, pie, la []bool, cnt, n, delta int) {
	if cnt >= n {
		*res++
		return
	}
	for i := 0; i < n; i++ {
		if cols[i] || pie[i+cnt] || la[i-cnt+delta] {
			continue
		}
		cols[i] = true
		pie[i+cnt] = true
		la[i-cnt+delta] = true
		dfs(res, cols, pie, la, cnt+1, n, delta)
		cols[i] = false
		pie[i+cnt] = false
		la[i-cnt+delta] = false
	}
}
func totalNQueens(n int) int {
	cols := make([]bool, n)
	pie := make([]bool, 2*n-1)
	la := make([]bool, 2*n-1)
	res := 0
	delta := n - 1
	dfs(&res, cols, pie, la, 0, n, delta)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 23:50:32
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
//
//
//
// Related Topics 回溯算法
// 👍 244 👎 0
