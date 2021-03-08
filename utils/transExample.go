// +ignore
package utils

import (
	"fmt"
	"strconv"
	"strings"
)

//func main() {
//	root := Trans2Tree("[3,9,20,null,null,15,7]")
//	PreorderTraversal(root, true)
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Trans2Tree(a string) *TreeNode {
	a = a[1:len(a)-1]
	arr := strings.Split(a, ",")
	if len(arr) < 1 || arr[0] == "null" {
		return nil
	}
	val, err := strconv.Atoi(arr[0])
	if err != nil {
		panic("atoi panic")
	}
	root := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	var tmpNode *TreeNode
	var tmpInt int
	breadthList := []*TreeNode{root}
	bakList := make([]*TreeNode, 0)
	left := true  // 应加入左节点?
	for _, v := range arr[1:] {
		if v == "null" {
			tmpNode = nil
		}else {
			tmpInt, _ = strconv.Atoi(v)
			tmpNode = &TreeNode{Val: tmpInt}
		}
		if len(breadthList) < 1 {
			breadthList = bakList
			bakList = make([]*TreeNode, 0)
			if len(breadthList) < 1 {
				return root
			}
		}
		if left {
			breadthList[0].Left = tmpNode
		}else {
			breadthList[0].Right = tmpNode
			breadthList = breadthList[1:]
		}

		left = !left
		if tmpNode != nil {
			bakList = append(bakList, tmpNode)
		}

	}
	return root
}


func PreorderTraversal(root *TreeNode, a... bool) {
	var printNull bool
	if len(a) > 0 {
		printNull = a[0]
	}

	if root != nil {
		fmt.Printf("%v, ", root.Val)
		PreorderTraversal(root.Left, printNull)
		PreorderTraversal(root.Right, printNull)
	}else {
		if printNull{
			fmt.Printf("null, ")
		}

	}
}

func LevelOrderTraversal(root *TreeNode, a... bool) {}




func Trans2Array(a string) []int {
	a = a[1:len(a)-1]
	arr := strings.Split(a, ",")
	res := make([]int, 0)
	intV := 0
	for _, v := range arr {
		intV, _ = strconv.Atoi(v)
		res = append(res, intV)
	}
	return res
}
