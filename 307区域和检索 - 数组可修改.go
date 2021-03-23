// +ignore

package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func (n *NumArray) sumRangeHelper(l, r, rangeL, rangeR, index int) int {
	if rangeL == l && rangeR == r {
		return n.tree[index]
	}
	mid := (r-l)/2 + l
	if rangeL > mid {
		return n.sumRangeHelper(mid+1, r, rangeL, rangeR, n.rightChild(index))
	} else if rangeR <= mid {
		return n.sumRangeHelper(l, mid, rangeL, rangeR, n.leftChild(index))
	} else {
		return merge(n.sumRangeHelper(l, mid, rangeL, mid, n.leftChild(index)),
			n.sumRangeHelper(mid+1, r, mid+1, rangeR, n.rightChild(index)))
	}
}

func (n *NumArray) maintain(l, r, index, target, value int) int {
	if l == r {
		n.tree[index] = value
		return value
	}
	mid := (r-l)/2 + l

	// todo: 这里的大于号需要确认一下
	if target > mid {
		n.tree[index] = merge(n.tree[n.leftChild(index)], n.maintain(mid+1, r, n.rightChild(index), target, value))
	} else {
		n.tree[index] = merge(n.tree[n.rightChild(index)], n.maintain(l, mid, n.leftChild(index), target, value))
	}
	return n.tree[index]
}

func merge(lValue, rValue int) int {
	// 这里决定了数据的整合方式, 两个区间是直接求和还是求差求余...
	return lValue + rValue
}

func (n *NumArray) buildSegementTree(index, l, r int) int {
	// 在index这个位置创建表示区间[l, r]的线段树, 即这里只是创建n.tree的一个index
	// 本来想在该函数直接return
	if l == r {
		n.tree[index] = n.data[l]
		return n.data[l]
	}
	n.tree[index] = merge(n.buildSegementTree(n.leftChild(index), l, (r-l)>>1+l),
		n.buildSegementTree(n.rightChild(index), (r-l)>>1+l+1, r))
	return n.tree[index]
}

func (n *NumArray) get(index int) int {
	if index < 0 || index >= len(n.data) {
		panic("index is illegal")
	}
	return n.data[index]
}

func (n *NumArray) rightChild(index int) int {
	return index<<1 + 2
}

type NumArray struct {
	data []int
	tree []int
}

func (n *NumArray) leftChild(index int) int {
	return index<<1 + 1
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tmp := n & (n - 1)
	if tmp != 0 {
		for tmp != 0 {
			n = tmp
			tmp = n & (n - 1)
		}
		n <<= 1
	}
	segementTree := make([]int, n<<1)
	res := NumArray{data: nums, tree: segementTree}
	res.buildSegementTree(0, 0, len(nums)-1)
	return res
}

func (this *NumArray) Update(index int, val int) {
	this.maintain(0, len(this.data)-1, 0, index, val)
	this.data[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	//if left == right {
	//	return this.data[left]
	//}
	return this.sumRangeHelper(0, len(this.data)-1, left, right, 0)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-22 14:34:06
//给你一个数组 nums ，请你完成两类查询，其中一类查询要求更新数组下标对应的值，另一类查询要求返回数组中某个范围内元素的总和。
//
// 实现 NumArray 类：
//
//
//
//
// NumArray(int[] nums) 用整数数组 nums 初始化对象
// void update(int index, int val) 将 nums[index] 的值更新为 val
// int sumRange(int left, int right) 返回子数组 nums[left, right] 的总和（即，nums[left] +
//nums[left + 1], ..., nums[right]）
//
//
//
//
// 示例：
//
//
//输入：
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//输出：
//[null, 9, null, 8]
//
//解释：
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // 返回 9 ，sum([1,3,5]) = 9
//numArray.update(1, 2);   // nums = [1,2,5]
//numArray.sumRange(0, 2); // 返回 8 ，sum([1,2,5]) = 8
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 104
// -100 <= nums[i] <= 100
// 0 <= index < nums.length
// -100 <= val <= 100
// 0 <= left <= right < nums.length
// 最多调用 3 * 104 次 update 和 sumRange 方法
//
//
//
// Related Topics 树状数组 线段树
// 👍 245 👎 0
