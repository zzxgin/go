/*
CSP  -->不要以共享内存的方式来通信，相反，要通过通信来共享内存

channel 本身就是一个队列  先进先出

channel 是 引用类型  可以使用make初始化   make--map、slice、channel
channel 写满的时候不能写， 取空的时候不能取  很容易发生阻塞 很容易产生死锁
每个channel 只能放同一种类型
使用channel 收发操作都是在不同的goroutine

channel

	无缓冲通道
	有缓冲通道
*/
package main

import (
	"fmt"
	"time"
)

// type signal struct{}

// func main() {
// 	ch := make(chan signal)
// 	go senddata(ch)
// 	fmt.Println("主协程正在等待数据...")
// 	<-ch // 主协程阻塞，直到收到信号
// 	fmt.Println("主协程继续执行")
// }

// func senddata(ch chan signal) {
// 	fmt.Println("子协程正在发送数据...")
// 	time.Sleep(2 * time.Second) // 模拟发送数据的耗时操作
// 	ch <- signal{}
// }

// // 通道的关闭  生产者发送完毕数据，关闭通道。消费者就不会因为没有数据产生死锁
// // 生产者
// func senddata(ch chan string) {
// 	defer close(ch) // 关键：函数退出前自动关闭通道，确保消费者能感知结束
// 	for i := 0; i < 3; i++ {
// 		ch <- fmt.Sprintf("发送数据:%d", i)
// 	}
// }

// // 消费者
// func main() {
// 	ch := make(chan string)
// 	go senddata(ch)

// 	// 方式一：通过ok变量判断通道是否关闭（二选一即可）
// 	/*
// 		for {
// 			data, ok := <-ch
// 			if !ok { // 通道关闭且无数据时，ok=false
// 				break
// 			}
// 			fmt.Println("接收数据（ok判断）：", data)
// 		}
// 		fmt.Println("所有数据接收完毕，程序正常退出")
// 	*/

// 	// 方式二：range 自动判断通道关闭（推荐，更简洁）
// 	for value := range ch {
// 		fmt.Println("接收数据（range遍历）：", value)
// 	}
// 	fmt.Println("所有数据接收完毕，程序正常退出")
// }

//有缓冲通道  发送数据时，如果缓冲区满了，发送者会被阻塞；接收数据时，如果缓冲区空了，接收者会被阻塞
// func senddata(ch chan string) {
// 	for i := 0; i < 5; i++ {
// 		ch <- fmt.Sprintf("data:%d", i)
// 		fmt.Println("往通道里放数据:", i)
// 	}
// 	defer close(ch) // 发送完所有数据后关闭通道（关键）
// }

// func main() {
// 	ch1 := make(chan string, 3) // 容量3的有缓冲通道
// 	go senddata(ch1)

// 	// 直接读取，依靠通道阻塞机制，无需 sleep
// 	for data := range ch1 {
// 		fmt.Println("读取通道数据:", data)
// 	}
// 	fmt.Println("所有数据读取完毕，程序正常退出")
// }

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// 启动生产者goroutine，发送完数据后关闭通道（关键）
	go func() {
		time.Sleep(10 * time.Second)
		ch1 <- 1
		close(ch1) // 发送完关闭
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
		close(ch2) // 发送完关闭
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch3 <- 3
		close(ch3) // 发送完关闭
	}()

	// 用计数器记录已接收的有效数据数，替代固定循环次数
	receivedCount := 0
	// 总数据量：3个通道各1条
	totalData := 3

	// 循环接收，直到所有数据都被读取
	for receivedCount < totalData {
		select {
		case msg, ok := <-ch1: // 检测通道是否关闭
			if ok {
				fmt.Println("接收到数据from ch1:", msg)
				receivedCount++
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Println("接收到数据from ch2:", msg)
				receivedCount++
			}
		case msg, ok := <-ch3:
			if ok {
				fmt.Println("接收到数据from ch3:", msg)
				receivedCount++
			}
		case <-time.After(15 * time.Second): // 超时时间覆盖最长发送时间（10s）
			fmt.Println("超时！部分数据未接收完成")
			return // 超时直接退出
		}
	}

	fmt.Println("所有数据接收完毕，程序正常退出")
}
