// +ignore
package main

func main() {
	findMedianSortedArrays([]int{1, 3}, []int{2, })
}

func dropLeft(nums1, nums2 []int, K int) int {
	if len(nums1) < 1 {
		return nums2[K - 1]
	}
	if len(nums2) < 1 {
		return nums1[K - 1]
	}

	if K == 1 {
		if nums1[0] > nums2[0] {
			return nums2[0]
		}
		return nums1[0]
	}
	var shortNums, longNums []int
	var shortSize int
	if len(nums1) > len(nums2) {
		shortNums = nums2
		longNums = nums1
		shortSize = len(nums2)
	}else {
		shortNums = nums1
		longNums = nums2
		shortSize = len(nums1)
	}
	halfK := K >> 1
	sub := halfK - 1
	if halfK > shortSize {
		sub = shortSize - 1
	}
	if longNums[sub] > shortNums[sub] {
		return dropLeft(longNums, shortNums[sub+1:], K - sub - 1)
	}else {
		return dropLeft(longNums[sub+1:], shortNums, K - sub - 1)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	if size & 1 == 1 {
		return float64(dropLeft(nums1, nums2, size - (size>>1)))
	}
	return (float64(dropLeft(nums1, nums2, size - (size>>1))) + float64(dropLeft(nums1, nums2, (size>>1) + 1))) / 2
}
