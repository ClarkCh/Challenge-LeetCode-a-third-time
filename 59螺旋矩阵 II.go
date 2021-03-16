// +ignore
package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func helper(reference [][]bool, x, y, n, funcNum int) (int, int, int) {
	switch funcNum & 3 {
	case 0:
		if x+1 < n && !reference[y][x+1] {
			return x + 1, y, funcNum
		}
		funcNum++
		fallthrough
	case 1:
		if y+1 < n && !reference[y+1][x] {
			return x, y + 1, funcNum
		}
		funcNum++
		fallthrough
	case 2:
		if x-1 >= 0 && !reference[y][x-1] {
			return x - 1, y, funcNum
		}
		funcNum++
		fallthrough
	case 3:
		if y-1 >= 0 && !reference[y-1][x] {
			return x, y - 1, funcNum
		}
		funcNum++
		if x+1 < n && !reference[y][x+1] {
			return x + 1, y, funcNum
		}
		funcNum++
	}
	return 0, 0, 0
}

func generateMatrix(n int) [][]int {
	reference := make([][]bool, n)
	for i := range reference {
		reference[i] = make([]bool, n)
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	x, y := 0, 0
	funcNum := 0
	for i := 1; i <= n*n; i++ {
		res[y][x] = i
		reference[y][x] = true
		x, y, funcNum = helper(reference, x, y, n, funcNum)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-16 09:27:56
//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
// Related Topics 数组
// 👍 334 👎 0
