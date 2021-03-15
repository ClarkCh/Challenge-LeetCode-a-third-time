// +ignore
package main

//func main() {
//	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6}, {7,8,9}}))
//	fmt.Println(spiralOrder([][]int{{6}, {9}, {7}}))
//}

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return []int{}
	}
	row, col := len(matrix), len(matrix[0])
	res := make([]int, 0)
	upRow, downRow, upCol, downCol := row-1, 0, col-1, 0
	i, j := 0, 0
	for {
		if upCol < downCol || upRow < downRow {
			break
		}
		for i <= upCol {
			res = append(res, matrix[j][i])
			i++
		}
		i--
		j++
		if j > upRow {
			break
		}
		for j <= upRow {
			res = append(res, matrix[j][i])
			j++
		}
		j--
		i--
		if i < downCol {
			break
		}
		for i > downCol {
			res = append(res, matrix[j][i])
			i--
		}
		for j > downRow {
			res = append(res, matrix[j][i])
			j--
		}
		j++
		i++

		upRow--
		upCol--
		downRow++
		downCol++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 09:53:10
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
// Related Topics 数组
// 👍 652 👎 0
