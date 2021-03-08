// +ignore
package main

import "fmt"

func main() {
	//fmt.Println(maxEnvelopes([][]int{{2, 3}, {3, 4}, {1, 2}, {3, 3}, {4, 7}}))
	//fmt.Println(maxEnvelopes([][]int{{5,4}, {6,4},{6,7},{2,3}}))
	//fmt.Println(maxEnvelopes([][]int{{46,89},{50,53},{52,68},{72,45}, {77,81}}))
	fmt.Println(maxEnvelopes([][]int{{7,8},{12,16},{12,5},{1,8},{4,19},{3,15},{4,10},{9,16}}))


}

//leetcode submit region begin(Prohibit modification and deletion)
type arrayList struct {
	Cnt int
	Val []int
	Next *arrayList
}
func add(a, b *arrayList) *arrayList {
	vHead := &arrayList{}
	vHead.Next = a
	prev := vHead
	for {
		if a.Val[1] > b.Val[1] {
			prev.Next = b
			b.Next = a
			break
		}else {
			if a.Next == nil {
				a.Next = b
				break
			}else {
				prev = a
				a = a.Next
			}

		}
	}

	return vHead.Next
}

func maxEnvelopes(envelopes [][]int) int {
	up := 20000000
	xArray := make([]*arrayList, up)
	for _, v := range envelopes {
		if xArray[v[0]] == nil {
			xArray[v[0]] = &arrayList{Val: v}
		}else {
			xArray[v[0]] = add(xArray[v[0]], &arrayList{Val: v})
		}
	}
	vHead := &arrayList{}
	prev := vHead
	for _, v := range xArray {
		if v != nil {
			prev.Next = v
			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}
	cur := vHead.Next
	cur.Cnt = 1
	cur = cur.Next
	res := 1
	for cur != nil {
		travel := vHead.Next
		for travel != cur {
			if cur.Val[0] > travel.Val[0] && cur.Val[1] > travel.Val[1] {
				if travel.Cnt >= cur.Cnt {
					cur.Cnt = travel.Cnt + 1
					if cur.Cnt > res {
						res = cur.Cnt
					}
				}
			}
			travel = travel.Next
		}
		if cur.Cnt < 1 {
			cur.Cnt = 1
		}
		cur = cur.Next
	}
	return res
}
//func maxEnvelopes(envelopes [][]int) int {
//	xMap := make(map[int][][]int)
//	yMap := make(map[int][][]int)
//	xArray := make([]int, 1000)
//	yArray := make([]int, 1000)
//	cntMap := make(map[int][][]int)
//	up := 100000
//	minX, minY := up, up
//	cnt := 1
//	for _, v := range envelopes {
//		x := v[0]
//		y := v[1]
//		xArray[x]++
//		yArray[y]++
//		if x < minX {
//			minX = x
//		}
//		if y < minY {
//			minY = y
//		}
//		vX, ok := xMap[x]
//		if !ok {
//			vX =  make([][]int, 0)
//		}
//		vX = append(vX, v)
//		xMap[x] = vX
//		vY, ok := yMap[y]
//		if !ok {
//			vY = make([][]int, 0)
//		}
//		vY = append(vY, v)
//		yMap[y] = vY
//	}
//	cntMap[cnt] = make([][]int, 0)
//	for _, v := range xMap[minX] {
//		cntMap[cnt] = append(cntMap[cnt], v)
//	}
//	for _, v := range yMap[minY] {
//		cntMap[cnt] = append(cntMap[cnt], v)
//	}
//	cnt++
//
//	x, y := 0, 0
//
//	for {
//		cntMap[cnt] = make([][]int, 0)
//		for i := x+1; i < up; i++ {
//			if xArray[i] > 0 {
//				x = i
//				goto label1
//			}
//		}
//		x = up
//		goto label2
//	label1:
//		for _, v := range xMap[x] {
//			for _, v2 := range cntMap[cnt-1] {
//				if v[0] > v2[0] && v[1] > v2[1] {
//					cntMap[cnt] = append(cntMap[cnt], v)
//				}
//			}
//		}
//	label2:
//		for i := y+1; i < up; i++ {
//			if yArray[i] > 0 {
//				y = i
//				goto label3
//			}
//		}
//		if up == x {
//			return cnt
//		}
//		y = up
//		goto label4
//	label3:
//		for _, v := range xMap[y] {
//			for _, v2 := range cntMap[cnt-1] {
//				if v[0] > v2[0] && v[1] > v2[1] {
//					cntMap[cnt] = append(cntMap[cnt], v)
//				}
//			}
//		}
//
//	label4:
//		cnt++
//	}
//
//}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-04 09:49:28[[5,4],[6,4],[6,7],[2,3]]
//给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如
//同俄罗斯套娃一样。 
//
// 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。 
//
// 说明: 
//不允许旋转信封。 
//
// 示例: 
//
// 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出: 3 
//解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
// 
// Related Topics 二分查找 动态规划 
// 👍 328 👎 0


