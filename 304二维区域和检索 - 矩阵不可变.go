// +ignore
package main

func main() {
	t := Constructor([][]int{{-1}})
	t.SumRegion(0, 0, 0, 0)
}
//leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	weight := len(matrix)
	if weight < 1 {
		return NumMatrix{[][]int{}}
	}
	height := len(matrix[0])

	res := make([][]int, weight)
	for i := 0; i < weight; i++ {
		res[i] = make([]int, height)
	}
	res[0][0] = matrix[0][0]
	for i := 1; i < weight; i++ {
		res[i][0] = res[i-1][0] + matrix[i][0]
	}
	for i := 1; i < height; i++ {
		res[0][i] = res[0][i-1] + matrix[0][i]
	}
	for i := 1; i < weight; i++ {
		for j := 1; j < height; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1] + matrix[i][j] - res[i-1][j-1]
		}
	}
	return NumMatrix{res}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	tmp1 := 0
	tmp2 := this.matrix[row2][col2]
	tmp3 := 0
	tmp4 := 0
	if row1 > 0 {
		if col1 > 0 {
			tmp1 = this.matrix[row1-1][col1-1]
		}
		tmp4 = this.matrix[row1-1][col2]
	}
	if col1 > 0 {
		tmp3 = this.matrix[row2][col1-1]
	}
	return tmp1 + tmp2 - tmp3 - tmp4
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-02 23:51:16
//给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。 
//
// 
//上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。 
//
// 
//
// 示例： 
//
// 
//给定 matrix = [
//  [3, 0, 1, 4, 2],
//  [5, 6, 3, 2, 1],
//  [1, 2, 0, 1, 5],
//  [4, 1, 0, 1, 7],
//  [1, 0, 3, 0, 5]
//]
//
//sumRegion(2, 1, 4, 3) -> 8
//sumRegion(1, 1, 2, 2) -> 11
//sumRegion(1, 2, 2, 4) -> 12
// 
//
// 
//
// 提示： 
//
// 
// 你可以假设矩阵不可变。 
// 会多次调用 sumRegion 方法。 
// 你可以假设 row1 ≤ row2 且 col1 ≤ col2 。 
// 
// Related Topics 动态规划 
// 👍 230 👎 0
