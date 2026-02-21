package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello")
	//GOMAXPROCS   设置当前P的数量  输入0表示查询当前P的数量，输入大于0的数，表示设置P的数量
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(8)
	fmt.Println(runtime.GOMAXPROCS(0))
	go hello()
	go hello2()
	//默认情况下，主协程退出，子协程也会退出,让主协程阻塞 2 秒，避免主协程提前退出。
	time.Sleep(2 * time.Second)
}
func hello() {
	fmt.Println("hello start...")
	time.Sleep(1 * time.Second)
	fmt.Println("hello end....")

}
func hello2() {
	fmt.Println("hello2 start...")
	time.Sleep(1 * time.Second)
	fmt.Println("hello2 end....")
}
