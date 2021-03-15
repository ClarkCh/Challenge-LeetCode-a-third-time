// +ignore
package main

func main() {
	//a := "baseljselfk"
	//fmt.Println(quickSortString(a))
	//fmt.Println(groupAnagrams([]string{"stop","pots","reed","","tops","deer","opts",""}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func quickSort(a []byte) {
	size := len(a)
	if size < 2 {
		return
	}
	l, move, r := 0, 1, size-1
	tmp := a[0]
	for move <= r {
		if a[move] < tmp {
			a[l], a[move] = a[move], tmp
			l++
			move++
		} else if a[move] == tmp {
			move++
		} else {
			a[r], a[move] = a[move], a[r]
			r--
		}
	}
	quickSort(a[:l])
	//if move < size - 1 {
	//	quickSort(a[move+1:])
	//}
	quickSort(a[r+1:])

}

func quickSortString(a string) string {
	size := len(a)
	res := []byte(a)
	if size < 2 {
		return a
	}

	quickSort(res)
	return string(res)
}

func groupAnagrams(strs []string) [][]string {
	size := len(strs)
	if size < 2 {
		return [][]string{strs}
	}
	rec := make(map[string][]string)
	for _, v := range strs {
		tmp := quickSortString(v)
		if v2, ok := rec[tmp]; ok {
			rec[tmp] = append(v2, v)
		} else {
			rec[tmp] = make([]string, 1, 1)
			rec[tmp][0] = v
		}
	}
	res := make([][]string, 0)
	for _, v := range rec {
		res = append(res, v)
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-03-15 14:52:31
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 681 👎 0
