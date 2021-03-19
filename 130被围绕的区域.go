// +ignore
package main

import "fmt"

func main() {
	board := [][]byte{{"X"[0], "X"[0], "X"[0], "X"[0]}, {"X"[0], "O"[0], "O"[0], "X"[0]},
		{"X"[0], "X"[0], "O"[0], "X"[0]}, {"X"[0], "O"[0], "X"[0], "X"[0]}}
	solve(board)
	fmt.Println(board)
}

//leetcode submit region begin(Prohibit modification and deletion)
func (u *UnionFind) union(a, b []int) {
	pRoot := u.find(a)
	qRoot := u.find(b)
	if pRoot == qRoot {
		return
	}
	if u.rank[pRoot] > u.rank[qRoot] {
		u.parents[qRoot] = pRoot
	} else if u.rank[pRoot] < u.rank[qRoot] {
		u.parents[pRoot] = qRoot
	} else {
		u.parents[pRoot] = qRoot
		u.rank[qRoot]++
	}
}

func (u *UnionFind) isConnected(a, b []int) bool {
	return u.find(a) == u.find(b)
}

func (u *UnionFind) array2int(a []int) int {
	return u.col*a[0] + a[1]
}

func (u *UnionFind) find(a []int) int {
	p := u.array2int(a)
	for p != u.parents[p] {
		u.parents[p] = u.parents[u.parents[p]]
		p = u.parents[p]
	}
	return p
}

type UnionFind struct {
	parents []int
	rank    []int
	col     int // y
}

func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	row := len(board)
	col := len(board[0])
	parents := make([]int, row*col)
	for i := 0; i < row*col; i++ {
		parents[i] = i
	}
	u := &UnionFind{
		parents: parents,
		rank:    make([]int, row*col),
		col:     col,
	}
	_ = u
	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			if board[j][i] == 79 {
				if board[j-1][i] == 79 {
					u.union([]int{j, i}, []int{j - 1, i})
				}
				if board[j][i-1] == 79 {
					u.union([]int{j, i}, []int{j, i - 1})
				}
			}
		}
	}

	edgeList := make([][]int, 0)
	tmpCol := col - 1
	tmpRow := row - 1
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if i == 0 || i == tmpCol || j == 0 || j == tmpRow {
				if board[j][i] == 79 {
					edgeList = append(edgeList, []int{j, i})
				}

			}
		}
	}

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if board[j][i] == 79 {
				for _, v := range edgeList {
					if u.isConnected([]int{j, i}, v) {
						goto label
					}
				}
				board[j][i] = 88
			}
		label:
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-17 17:56:20
//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
//
//
//
//
// 示例 1：
//
//
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X"
//,"X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//
//
// 示例 2：
//
//
//输入：board = [["X"]]
//输出：[["X"]]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
//
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 492 👎 0
