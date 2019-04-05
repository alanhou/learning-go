package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (aa = 3
	ss = "kkk"
	bb = true
	)
//函数外不能使用 := 来进行赋值

//定义变量
func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//变量赋初值
func variableInitialValue(){
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction(){
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShort(){
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler(){
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	// 与上面一句效果相同：fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%0.3f\n", cmplx.Exp(1i * math.Pi) + 1) //浮点数不精确，取0的方法
}

// 类型转换
func triangle(){
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量
func consts(){
	const (
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c = int(math.Sqrt(a*a + b*b)) // a,b 未指定类型无需转换为 float
	fmt.Println(filename, c)
}

// 枚举类型
func enums(){
	const(
		cpp = iota // 用于自增，可不再为下面的赋值
		java
		python
		golang
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
}
