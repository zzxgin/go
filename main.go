package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0700)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	} else {
		data := make([]byte, 100)
		for {
			len, err := file.Read(data)
			if len == 0 || err != nil {
				break
			}
			fmt.Println("每次读取: ", string(data))
		}
		defer file.Close() // 关闭文件，释放资源

	}
}

/***
go项目包管理方式：
1、gopath
早期的go项目环境是看gopath的设置，项目都要放在$GOPATH/src目录下， 下载的第三方包也放在$GOPATH/pkg目录下
这种管理会造成多个项目混用，管理起来不方便，也不好控制版本

2、go vender
项目还是放在$GOPATH目录下面，引入了一个vendor可以进行不同项目的版本管理

3、go mod
a、项目工程目录可以放在GOPATH路径之外
b、要求项目中必须要有go.mod文件，go.mod文件记录当前项目需要的第三方软件以及它的版本
c、所有第三方依赖包都放在$GOPATH/pkg/mod 目录下
***/
