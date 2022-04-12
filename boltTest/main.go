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
 * @DateTime: 2022/3/22 10:44
 **/
func main() {
	fmt.Println()
	/*
		使用终端编译程序：go build main.go
		终端运行：main.exe
		编译的时候,也可以传递参数进行
		编译的时候,也可以传递参数进行
	
	*/
	//接收运行时的参数,返回String[]切片数据
	args := os.Args

	fmt.Println("遍历args[]切片")
	//从切片拿数据  单个：下标拿  拿所有：循环
	for key, value := range args {
		fmt.Println(key,value)
	}

	//设定类型  可以设定默认值
	i := flag.Int("age", 18, "年龄")
	name := flag.String("name","linfeifei","名字")
	//                        参数名        参数默认值         参数描述

	//直接输出了默认值，这是不对的，应该编译解析

	var name1 string
	flag.StringVar(&name1,"name3","tom","名字")
	//          接收的变量      参数名       默认值       参数描述

	//编译解析
	flag.Parse()
	fmt.Println("编译解析后输出：")

	fmt.Println(*i)
	fmt.Println(*name)
	fmt.Println(name1)
}


func Menu() {

	fmt.Println("**********功能菜单*********")
	fmt.Println("1、创建区块链")
	fmt.Println("2、添加区块")
	fmt.Println("3、遍历区块")
	fmt.Println("4、查看区块")
	fmt.Println("5、查看帮助菜单")
	fmt.Println("请输入你的功能...")
	var num int
	fmt.Scanln(&num)
	fmt.Println(num)

	switch num {
	case 1:
		A()
	case 2:
		B()
	case 3:
		C()
	case 4:
		D()
	default:
		fmt.Println("输入有误，请检查")

	}

}
func A() {

}
func B() {

}
func C() {

}
func D() {

}