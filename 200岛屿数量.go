// +ignore
package main

func main() {
	//a := [][]byte{{"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"0"[0],"1"[0],"1"[0],"1"[0],"1"[0]},{"1"[0],"0"[0],"1"[0],"0"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0]},{"0"[0],"1"[0],"1"[0],"1"[0],"0"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0]},{"1"[0],"1"[0],"0"[0],"1"[0],"1"[0],"0"[0],"0"[0],"0"[0],"0"[0],"1"[0]},{"1"[0],"0"[0],"1"[0],"0"[0],"1"[0],"0"[0],"0"[0],"1"[0],"0"[0],"1"[0]},{"1"[0],"0"[0],"0"[0],"1"[0],"1"[0],"1"[0],"0"[0],"1"[0],"0"[0],"0"[0]},{"0"[0],"0"[0],"1"[0],"0"[0],"0"[0],"1"[0],"1"[0],"1"[0],"1"[0],"0"[0]},{"1"[0],"0"[0],"1"[0],"1"[0],"1"[0],"0"[0],"0"[0],"1"[0],"1"[0],"1"[0]},{"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"0"[0],"1"[0]},{"1"[0],"0"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"1"[0],"0"[0]}}
	//for i := 0; i < len(a); i++ {
	//	for j := 0; j < len(a[0]); j++ {
	//		fmt.Printf("%v, ", a[i][j] - 48)
	//	}
	//	fmt.Println()
	//}
	//fmt.Println(numIslands(a))
}

//leetcode submit region begin(Prohibit modification and deletion)
func (u *unionFind) union(a, b []int) {
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

func (u *unionFind) isConnected(a, b []int) bool {
	return u.find(a) == u.find(b)
}

func (u *unionFind) array2int(a []int) int {
	return a[1]*u.x + a[0]
}

func (u *unionFind) find(a []int) int {
	p := u.array2int(a)
	for p != u.parents[p] {
		u.parents[p] = u.parents[u.parents[p]]
		p = u.parents[p]
	}
	return p
}

type unionFind struct {
	parents []int
	rank    []int
	x       int
}

func numIslands(grid [][]byte) int {
	x, y := len(grid[0]), len(grid)
	parents := make([]int, x*y)
	rank := make([]int, x*y)
	for i := 0; i < x*y; i++ {
		parents[i] = i
	}
	u := &unionFind{
		parents: parents,
		rank:    rank,
		x:       x,
	}
	oneList := make([][]int, 0)
	if grid[0][0] == 49 {
		oneList = append(oneList, []int{0, 0})
	}
	for i := 1; i < x; i++ {
		if grid[0][i] == 49 {
			if grid[0][i-1] == 49 {
				u.union([]int{i - 1, 0}, []int{i, 0})
			}
			oneList = append(oneList, []int{i, 0})
		}
	}
	for i := 1; i < y; i++ {
		if grid[i][0] == 49 {
			if grid[i-1][0] == 49 {
				u.union([]int{0, i - 1}, []int{0, i})
			}
			oneList = append(oneList, []int{0, i})
		}
	}
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if grid[j][i] == 49 {
				if grid[j-1][i] == 49 {
					u.union([]int{i, j}, []int{i, j - 1})
				}
				if grid[j][i-1] == 49 {
					u.union([]int{i, j}, []int{i - 1, j})
				}
				oneList = append(oneList, []int{i, j})
			}
		}
	}

	cntMap := make(map[int]bool)
	for _, v := range oneList {
		p := u.find(v)
		cntMap[p] = true
	}
	return len(cntMap)
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-18 09:11:50
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 1038 👎 0
