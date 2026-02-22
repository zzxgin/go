package main // 声明当前代码属于main包，可独立运行

import (
	"encoding/json" // 导入JSON序列化/反序列化标准库
	"fmt"           // 导入格式化输出库
)

/*
这里的注释是对反射的说明，核心点：
1. 反射是程序运行时获取/修改自身信息的能力
2. Go的reflect包提供TypeOf/ValueOf获取变量的类型和值
3. 这段注释是铺垫——因为json.Marshal底层就是用反射实现的
*/

// 定义结构体StuInfo，用于存储学生信息
type StuInfo struct {
	// 结构体字段 + 标签（Tag）：反引号``包裹，是给程序读取的元信息
	StuId int    `json:"stu_id"` // json标签：序列化时字段名转为stu_id（小写+下划线）
	Name  string `json:"name"`   // json标签：序列化时字段名保持name（和结构体字段名一致）
	Age   int    `json:"-"`      // json标签：-表示序列化时忽略该字段（Age不会出现在JSON中）
}

func main() { // 程序入口函数
	// 1. 创建StuInfo结构体的指针实例，初始化字段值
	// StuId=1，Name="test"，Age=25
	stu := &StuInfo{1, "test", 25}

	// 2. 将结构体实例序列化为JSON字节切片（[]byte类型）
	// json.Marshal是核心函数：
	// - 底层通过反射遍历结构体的所有字段
	// - 读取每个字段的json标签，按标签规则处理
	// - 返回序列化后的字节切片，以及可能的错误（比如循环引用）
	stu_json, err := json.Marshal(stu)

	// 3. 错误处理：如果序列化失败（比如字段类型不支持JSON），打印错误
	if err != nil {
		fmt.Println(err)
	}

	// 4. 将JSON字节切片转为字符串并打印
	// 最终输出：{"stu_id":1,"name":"test"}（Age字段被忽略）
	fmt.Println(stu_json)
	fmt.Println(string(stu_json))
}
