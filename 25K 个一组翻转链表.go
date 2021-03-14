// +ignore
package main

import "fmt"

func main() {
	vHead := &ListNode{}
	pre := vHead
	for i := 0; i < 20; i++ {
		pre.Next = &ListNode{Val: i}
		pre = pre.Next
	}
	head := reverseKGroup(vHead.Next, 3)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
//leetcode submit region begin(Prohibit modification and deletion)
//func helper2(cur *ListNode, n, i int) (*ListNode, *ListNode, *ListNode, bool) {
//	if n == i {
//		nextHead := cur.Next
//		return cur, cur, nextHead, true
//	}
//	if cur != nil {
//		if cur.Next == nil {
//			return cur, nil, nil, false
//		}
//		head, tail, nextHead, ok := helper2(cur.Next, n, i+1)
//		if ok {
//			tail.Next = cur
//			return head, cur, nextHead, true
//		}
//		return cur, nil, nil, false
//	}else {
//		return cur, nil, nil, false
//	}
//}
//
//func helper(pre *ListNode, n int) *ListNode {
//	tmp, tail, nextHead, ok := helper2(pre.Next, n, 1)
//	pre.Next = tmp
//	tail.Next = nextHead
//	if ok && nextHead != nil {
//		return helper(tail, n)
//	}
//	return tmp
//}

func helper(cur *ListNode, i, k int) (*ListNode, *ListNode) {

	if k == i {
		return cur, cur
	}
	newHead, pre := helper(cur.Next, i+1, k)
	pre.Next = cur
	return newHead, cur
}

//type ListNode struct {
//    Val int
//    Next *ListNode
//}
func reverseKGroup(head *ListNode, k int) *ListNode {
	vHead := &ListNode{}
	vHead.Next = head
	pre := vHead
	//return helper(vHead, k)
	cnt := 0
	inputHead := pre
	for pre.Next != nil {
		cnt++
		pre = pre.Next
		if cnt == k {
			nextHead := pre.Next
			newHead, newTail := helper(inputHead.Next, 1, k)
			inputHead.Next = newHead
			newTail.Next = nextHead
			pre = newTail
			inputHead = pre
			cnt = 0
		}


	}
	return vHead.Next
}
//leetcode submit region end(Prohibit modification and deletion)


// 2021-03-07 16:26:38
//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。 
//
// k 是一个正整数，它的值小于或等于链表的长度。 
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 
//
// 进阶： 
//
// 
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？ 
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
// 
//
// 示例 2： 
//
// 
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
// 
//
// 示例 3： 
//
// 
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
// 
//
// 示例 4： 
//
// 
//输入：head = [1], k = 1
//输出：[1]
// 
//
// 
// 
//
// 提示： 
//
// 
// 列表中节点的数量在范围 sz 内 
// 1 <= sz <= 5000 
// 0 <= Node.val <= 1000 
// 1 <= k <= sz 
// 
// Related Topics 链表 
// 👍 940 👎 0
