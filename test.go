// +ignore
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum2([]int{-1,0,1,2,-1,-4,-2,-3,3,0,4}))
}

func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	res := make([][]int, 0)
	last := nums[0] - 1
	for mid := 0; mid < len(nums) - 2; mid++ {
		if last == nums[mid] {
			continue
		}
		l = mid + 1
		r = len(nums) - 1
		tmp := - nums[mid]
		leftHalfTmp := tmp / 2
		rightHalfTmp := tmp - leftHalfTmp
		smallLast := last - 1

		for l < r {
			if smallLast == nums[l] {
				l++
				continue
			}
			if smallLast == nums[r] {
				r--
				continue
			}
			//if nums[l] * 2 > tmp || nums[r] * 2 < tmp {
			//	goto label
			//}
			if nums[l] > leftHalfTmp || nums[r] < rightHalfTmp {
				goto label
			}
			if nums[l] + nums[r] > tmp {
				smallLast = nums[r]
				r--

			}else if nums[l] + nums[r] < tmp {
				smallLast = nums[l]
				l++

			}else {
				res = append(res, []int{nums[mid], nums[l], nums[r]})
				smallLast = nums[r]
				r--
			}
		}
	label:
		last = nums[mid]
	}
	return res
}
