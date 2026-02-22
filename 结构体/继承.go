package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
	Age  int
	Sex  string
}

type Teacher struct {
	Person
	Person2 // 继承多个结构体
	Subject string
	//Name    int // 重写字段（放开注释会覆盖Person的Name）
}

// Person的方法，会被Teacher继承
func (p *Person) Hello() {
	fmt.Println("我是Person")
}

func main() {
	// 方式1：直接实例化Teacher
	t := Teacher{
		Person:  Person{Name: "张三", Age: 30},
		Subject: "数学",
	}
	fmt.Printf("Teacher信息：%+v\n", t)
	// 方式2：先实例化Person，再通过Person来初始化Teacher

	p := Person{Name: "李四", Age: 40}
	t1 := Teacher{
		Person:  p,
		Subject: "英语",
	}
	fmt.Printf("Teacher信息：%+v\n", t1)
	t.Hello()        // 隐式调用（方法提升）
	t.Person.Hello() // 显式调用

	t.Person.Name = "李四" //Teacher继承多个结构体时，属性字段如果发生冲突，必须分开写
	t.Person2.Name = "王五"
}
