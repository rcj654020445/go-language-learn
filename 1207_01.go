package main

import "fmt"

//第一种，指定变量类型
var age int = 18
//第二种，根据值自行判定变量类型
var name  = "菜鸟教程"


/*
多变量声明
*/

//和python很像,不需要显示声明类型，自动推断
var c,d = "c",12

//因式分解关键字的写法一般用于声明全局变量
var (
 e int
 f string
)



func main() {
	//类型相同多个变量, 非全局变量
	var a,b int
	a,b = 10,11
	//这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	//第三种,省略var
    company := "微众"
	fmt.Println(company)
}