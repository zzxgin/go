package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到异常：", err)
		}
	}()
	fmt.Println("这是test函数")
	lst := []int{1, 2, 3}
	fmt.Println(lst[3]) // 访问越界，产生异常
}
func test2() {
	fmt.Println("这是test2函数")
	time.Sleep(1 * time.Second)
	fmt.Println("test2函数结束")
}
func main() {
	fmt.Println("这是main函数")
	go test()
	go test2()
	time.Sleep(2 * time.Second)
	fmt.Println("main函数结束")
}
