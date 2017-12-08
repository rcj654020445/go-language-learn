package main
import "fmt"
import "unsafe"



func main() {
	const LENGTH int = 10

	const WIDTH  int = 5;

	var area int;

	const A,B,C = 1,false,"str"//多重赋值
	area = LENGTH * WIDTH

	fmt.Printf("面积为：%d",area);
	fmt.Println()
	fmt.Println(A,B,C)


	/*
	常量还可以用作枚举;
	用函数计算表达式的值作为常量的值
	*/
	const (
        Unknown = 0
        Femal = 1
        Male = 2
        Str1 = "abc"
        Str2 = len(Str1)//用len()函数计算表达式的值
        Str3 = unsafe.Sizeof(Str1)//用unsafe.Sizeof()函数计算表达式的值
	)
	// *****note******: 常量表达式中，函数必须是内置函数



    //iota
    const (
       Str4 = iota
       Str5 = iota
    )
    Str4
    Str5
    Str4
    Str5
    fmt.Println(Str4,Str5)
}
