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
 * @DateTime: 2022/3/29 9:15
 **/

func main() {
	//获取用户输入的命令行参数
	args := os.Args
	//接收输入的参数集合
	today := flag.NewFlagSet("today", flag.ExitOnError)
	//功能参数   遇见错误处理方式,退出

	//通过下标判断输入功能的是否为today
	fmt.Println(args[1]);
	if args[1] == "today" {
		//输入today功能的请求参数
		day := today.String("day","星期一","星期")
		//输入sun功能的请求参数
		sun := today.String("sun","晴天","天气")

		//输出使用前需要编译解析
		today.Parse(args[2:])  //获取today功能后面的所有参数

		//获取参数，调用函数处理请求
		Today(*day)
		fmt.Println(*sun)
	}

	/*//功能解析，执行相应
	switch args[1] {

	}*/
}
func Today(day string) {
	fmt.Println(day)
}