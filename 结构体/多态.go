package main

import "fmt"

// 1. 定义接口：约定“支付”这个行为（只定规则，不实现）
type MoneyPay interface {
	Pay()
}

type WX struct {
	Name string
}
type ZFB struct {
	Name string
}

// 3. 给微信支付宝实现MoneyPay接口的pay()方法
// 这一步是关键：WX类型（准确说是*WX）实现了MoneyPay接口的所有方法（这里只有pay()）
// 因此*WX类型就“符合”MoneyPay接口的要求，可以被当作MoneyPay类型使用
func (x *WX) Pay() {
	fmt.Println("微信支付")
}

//同理
func (z *ZFB) Pay() {
	fmt.Println("支付宝支付")
}

//定义通用支付函数：接收MoneyPay接口类型参数
// 这个函数不关心传入的是支付宝还是微信，只要它实现了pay()方法（符合MoneyPay接口），就能调用pay()
// 这就是多态的核心：同一函数，处理不同类型，执行不同的pay()逻辑
func FinPay(p MoneyPay) {
	p.Pay() // 调用的是传入类型（ZFB/WX）各自实现的pay()方法
}

func main() {
	z := &ZFB{Name: "支付宝"}
	w := &WX{Name: "微信"}
	// 接口变量接收不同类型的实例
	// j1是MoneyPay类型，但可以存储*ZFB类型（因为*ZFB实现了MoneyPay接口）
	// j2是MoneyPay类型，但可以存储*WX类型（因为*WX实现了MoneyPay接口）
	var j1, j2 MoneyPay
	j1 = z
	j2 = w
	fmt.Printf("接口变量j1: %v, j2: %v\n", j1, j2)
	// 调用通用支付函数
	// 传入z（支付宝），FinPay会执行ZFB的pay()
	FinPay(z)
	// 传入w（微信），FinPay会执行WX的pay()
	FinPay(w)
}
