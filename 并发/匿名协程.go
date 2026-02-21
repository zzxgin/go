// 匿名协程
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i=", i)
			time.Sleep(1 * time.Second)
		}

	}()

	for i := 0; i < 3; i++ {
		fmt.Println("主协程 i=", i)
		time.Sleep(1 * time.Second)
	}
}
