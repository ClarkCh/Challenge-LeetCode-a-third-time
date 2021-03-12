// +ignore
package main

import "sort"

func main() {

}

// 参见39题, 但39题允许重用元素, 这里的是不可重用元素
type resStruct struct {
	candidates []int
	res        [][]int
	target     int
	size       int
}

func (r *resStruct) sumNotAddHelper(sumList []int, sum, index int) {
	if index >= r.size {
		return
	}
	temporary := sum + r.candidates[index]
	if temporary < r.target {
		r.sumNotAddHelper(sumList, sum, index+1)
	} else if temporary > r.target {
		return
	} else {
		r.res = append(r.res, append(sumList, r.candidates[index]))
	}
}

func (r *resStruct) helper(sumList []int, sum, index int) {
	if index >= r.size {
		return
	}
	temporary := sum + r.candidates[index]
	if temporary < r.target {
		r.helper(append(sumList, r.candidates[index]), temporary, index+1)
		//r.helper(sumList, sum, index+1)
		//r.sumNotAddHelper(append(sumList, r.candidates[index]), temporary, index+1)
	} else if temporary > r.target {
		return
	} else {
		r.res = append(r.res, append(sumList, r.candidates[index]))
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	size := len(candidates)
	for i, v := range candidates {
		if v > target {
			if size < i {
				size = i
			}
		}
	}
	res := &resStruct{
		candidates: candidates,
		res:        make([][]int, 0),
		target:     target,
		size:       size,
	}
	for i := 0; i < size; i++ {
		res.helper(make([]int, 0), 0, i)
		//res.sumNotAddHelper(make([]int, 0), 0, 0)
	}
	return res.res
}
