package main

import (
	"flag"
	"fmt"
	"os"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/29 9:43
 **/
/*
	简单计算器
*/
func main() {
	//声明os.args
	args := os.Args
	fmt.Println("功能有：add、sub、mul、div")

	//1、指定功能
	add := flag.NewFlagSet("add", flag.ExitOnError)
	sub := flag.NewFlagSet("sub", flag.ExitOnError)
	mul := flag.NewFlagSet("mul", flag.ExitOnError)
	div := flag.NewFlagSet("div", flag.ExitOnError)

	switch args[1] {
	case "add":
		//2、获取参数
		a := add.Float64("num1", 1, "num1")
		b := add.Float64("num2", 2, "num2")
		//3、编译解析 从指定位置开始:
		add.Parse(args[2:])
		//4、调用对应方法
		Add(*a, *b)
	case "sub":
		a := sub.Float64("num1", 1, "num1")
		b := sub.Float64("num2", 2, "num2")
		sub.Parse(args[2:])
		Sub(*a, *b)
	case "mul":
		a := mul.Float64("num1", 1, "num1")
		b := mul.Float64("num2", 2, "num2")
		mul.Parse(args[2:])
		Mul(*a, *b)
	case "div":
		a := div.Float64("num1", 1, "num1")
		b := div.Float64("num2", 2, "num2")
		div.Parse(args[2:])
		Div(*a, *b)
	default:
		fmt.Println("没有该功能~请检查！！！")
		//退出
		os.Exit(1)
	}

}

/*
	加减乘除功能
*/
func Add(a, b float64) {
	fmt.Println("add结果：", a+b)
	fmt.Println("林菲菲")
}

func Sub(a, b float64) {
	if a > b {
		fmt.Println("sub结果：", a-b)
		fmt.Println("林菲菲")

	} else if b > a {
		fmt.Println("sub结果：", b-a)
		fmt.Println("林菲菲")

	} else {
		fmt.Println("sub结果：", a-b)
		fmt.Println("林菲菲")

	}
}
func Mul(a, b float64) {
	fmt.Println("mul结果：", a*b)
	fmt.Println("林菲菲")

}

func Div(a, b float64) {
	if a > b {
		fmt.Println("div结果：", a/b)
		fmt.Println("林菲菲")

	} else if b > a {
		fmt.Println("div结果：", b/a)
		fmt.Println("林菲菲")

	} else { //相等
		fmt.Println("div结果：", a/b)
		fmt.Println("林菲菲")

	}
}
