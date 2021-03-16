// +ignore
package main

import (
	"fmt"
	"sort"
)

type uint8Struct []uint8

func (self uint8Struct) Len() int {
	return len(self)
}

func (self uint8Struct) Less(i, j int) bool {
	if self[i] < self[j] {
		return true
	}
	return false
}

func (self uint8Struct) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func main() {
	a := uint8Struct{3, 2, 4, 2, 8}
	sort.Sort(a)
	fmt.Println(a)
}

func QuickSort2(nums []int) {
	sort.Ints(nums)
}

func QuickSort(nums []int) {
	size := len(nums)
	if size < 2 {
		return
	}
	l, move, r := 0, 1, size-1
	tmp := nums[0]
	for move <= r {
		if nums[move] < tmp {
			nums[l] = nums[move]
			nums[move] = tmp
			l++
			move++
		} else if nums[move] == tmp {
			move++
		} else {
			nums[r], nums[move] = nums[move], nums[r]
			r--
		}
	}
	QuickSort(nums[:l])
	QuickSort(nums[r+1:])
}
