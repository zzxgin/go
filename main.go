package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x, y := 10, 2
	fmt.Println("输出：", x/y)
	fmt.Println(math.Pi)
	fmt.Println(5 | 4)

	if v, e := strconv.Atoi("abc"); e == nil {
		fmt.Println("转换成功：", v)
	} else {
		fmt.Println("转换失败：", e)
	}
	var day int
	fmt.Print("请输入一个数字：")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("周一")
	case 2, 3: // case多值匹配
		fmt.Println("周二/周三")
	default:
		fmt.Println("其他")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("循环：", i)
	}
	for {
		fmt.Println("循环")
		break
	}
	str := "Go语言"
	for idx, char := range str {
		fmt.Printf("索引：%d，字符：%c（Unicode值：%U）\n", idx, char, char)
	}
}
