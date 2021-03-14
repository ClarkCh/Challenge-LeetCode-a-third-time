// +ignore
package main

import "fmt"

func main() {
	a := []int{2}
	b := a[1:]
	fmt.Println(b)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
