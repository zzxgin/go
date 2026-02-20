package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	stu1 := new(Student)
	stu1.Name = "张三"
	stu1.Age = 20
	fmt.Printf("学生信息：%v\n", stu1.Name)

	stu2 := &Student{
		Name: "李四",
		Age:  22,
	}
	fmt.Printf("学生信息：%v\n", stu2.Name)

	stu3 := Student{
		Name: "王五",
		Age:  25,
	}
	fmt.Printf("学生信息：%v\n", stu3.Name)

}
