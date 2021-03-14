// +ignore
package main

import (
	"container/heap"
	"fmt"
	"unsafe"
)

//func main() {
	//a := &ListNode{Val: 2222222222}
	//fmt.Println(mergeKLists([]*ListNode{a}))
	//fmt.Println(mergeKLists(nil))
//}

//leetcode submit region begin(Prohibit modification and deletion)
func (l *ListNodeHeap) Pop() interface{} {
	size := len(*l)
	if size < 1 {
		return nil
	}
	res := (*l)[size-1]
	*l = (*l)[:size-1]
	return res
}

func (l *ListNodeHeap) Push(node interface{}) {
	*l = append(*l, node.(*ListNode))
}

func (l ListNodeHeap) Less(i, j int) bool {
	if l[i].Val < l[j].Val {
		return true
	}
	return false
}

func (l ListNodeHeap) Swap(i, j int) {
	if j < 0 {
		return
	}
	l[i], l[j] = l[j], l[i]
}

type ListNodeHeap []*ListNode

func (l ListNodeHeap) Len() int {
	return len(l)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil ||  len(lists) < 1 {
		return nil
	}
	//if len(lists) == 1 && lists[0] == nil {
	//	return nil
	//}
	nodeList := make([]*ListNode, 0)
	nodeHeap := (*ListNodeHeap)(unsafe.Pointer(&nodeList))
	heap.Init(nodeHeap)

	vHead := &ListNode{}
	pre := vHead
	for _, v := range lists {
		if v != nil {

			heap.Push(nodeHeap, v)
		}
	}
	for {
		tmp2 := heap.Pop(nodeHeap)
		if tmp2 == nil {
			return vHead.Next
		}
		tmp := tmp2.(*ListNode)
		pre.Next = tmp
		pre = pre.Next
		tmp = tmp.Next
		if tmp != nil {
			heap.Push(nodeHeap, tmp)
		}

	}

}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-07 03:04:01
//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 堆 链表 分治算法
// 👍 1183 👎 0
