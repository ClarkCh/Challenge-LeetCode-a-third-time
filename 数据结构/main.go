// +ignore
package main

import "fmt"

func main() {
	for i := 0; i < 209; i++ {
		test(i)
	}
}

func test(n int) {
	rawN := n
	//size := len(nums)
	tmp := n & (n - 1)
	if tmp != 0  {
		for tmp != 0 {
			n = tmp
			tmp = n & (n - 1)
		}
		n <<= 1
	}

	fmt.Printf("%v: %v\n", rawN, n)
}
