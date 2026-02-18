//go:build ignore

package main

import "fmt"

func main() {
	// 数组定义
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{2, 3, 4, 5, 6}
	for _, v := range arr1 {
		for _, w := range arr2 {
			if v == w {
				fmt.Printf("公共元素：%d\n", v)
			}
		}
	}
}
