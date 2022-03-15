package main

import "fmt"

func main() {
	/**
	1.2 五种创建变量的方法
	1.2.1 一行声明一个变量 var是关键字，name是变量名,type是类型,也可以不指定类型，Go会对其进行隐式初始化,string类型为空字符串,
	      float初始化为0,0 ，bool类型初始化为false,指针类型就初始化为nil
	 var name string = "Go编程时光"  多数情况下推荐指定类型，不用偷懒
	*/
	//var name = "Go编程时光"

	/**
	1.2.2 多个变量一起声明
	var (
		name   string = "string"
		age    int    = 0
		gender string = "女性"
	)
	*/
	//fmt.Println(name, age, gender)

	/*
		1.2.3 声明和初始化一个变量
		使用 := (推导声明写法或者短类型声明法:编译器会自动根据右值类型推断出左值的对应类型),可以声明一个变量,并对其进行(显式)初始化.
	*/
	//name := "Go编程时光" 等价于 var name string = "Go编程时光"
	// 等价于 var name  =  "Go编程时光"

	/**
	1.2.4 声明和初始化多个变量
	*/
	//name, age := "wangbm", 28 //这种方法，也经常用于变量的交换
	//var a int = 100
	//var b int = 200
	//b, a = a, b
	//fmt.Println(a, b)

	/**
	1.2.5 new函数声明一个指针变量 ，变量分为两种普通变量和指针变量,普通变量，存放的是数据本身,而指针变量存放的是数据的地址
	*/
	//如下代码，age是一个普通变量，存放的内容是28,而ptr是存放变量age值的内存地址:0xc000022098
	//var age int = 28
	//var prt = &age // &后面接变量名，表示取出该变量的内存地址
	//fmt.Println(prt, age)
	//输出 0xc000022098 28

	//而 这里说的new 函数，是go里的内建函数,使用表达式new(Type)将创建一个Type类型的匿名变量,初始化为Type类型的零值,然后返回变量地址,
	//返回的指针类型为*Type。

	//ptr := new(int)
	//fmt.Println("prt address: ", ptr)
	//fmt.Println("prt value: ", *ptr) //* 后面接指针变量,表示从内存地址中取出值
	//输出
	//prt address:  0xc000122008
	//prt value:  0

	//用new 创建变量和普通变量声明语句方式创建变量没有什么区别,除了不需要声明一个临时变量的没名字外，我们还可以在表达式中使用
	//new(Type).换言之,new 函数类似是一种语法糖,而不是一个新的基础概念。

	//使用new
	//func newInt() *int{
	//	return new(int)
	//}

	//使用传统的方式
	//func newInt() *int {
	//	var dummy int
	//	return &dummy
	//}
	//以上不管哪种方法,变量/常量都只能声明一次,声明多次,编译就会报错

	//特殊变量:也称占位符，或者空白标志符,用下划线表示.
	//匿名变量:优点有三:
	//1.不分配内存，不占用内存空间
	//2.不需要你为命名无用的变量名而纠结
	//3.多次声明不会有任何问题
	//通常我们用匿名接收必须接收，但是又不会用到的值

	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}
func GetData() (int, int) {
	return 100, 200
}
