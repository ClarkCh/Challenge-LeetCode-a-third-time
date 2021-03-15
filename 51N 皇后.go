// +ignore
package main

func main() {
	//fmt.Println(".Q"[0])
	//fmt.Println(".Q"[1])
}

//leetcode submit region begin(Prohibit modification and deletion)
func dfs(res *[][]string, rec [][]byte, cols, pie, la []bool, cnt, n, delta int) {
	if cnt >= n {
		tmp := make([]string, len(rec))
		for i := range tmp {
			tmp[i] = string(rec[i])
		}
		*res = append(*res, tmp)
	}
	for i := 0; i < n; i++ {
		if cols[i] || pie[i+cnt] || la[i-cnt+delta] {
			continue
		}
		cols[i] = true
		pie[i+cnt] = true
		la[i-cnt+delta] = true
		rec[cnt][i] = 81
		dfs(res, rec, cols, pie, la, cnt+1, n, delta)
		rec[cnt][i] = 46
		cols[i] = false
		pie[i+cnt] = false
		la[i-cnt+delta] = false
	}
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	cols := make([]bool, n)
	pie := make([]bool, 2*n-1)
	la := make([]bool, 2*n-1)
	res := make([][]string, 0)
	rec := make([][]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = make([]byte, n)
	}
	for _, v := range rec {
		for i := range v {
			v[i] = 46
		}
	}
	delta := n - 1
	dfs(&res, rec, cols, pie, la, 0, n, delta)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 17:30:24
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
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
// 👍 785 👎 0
