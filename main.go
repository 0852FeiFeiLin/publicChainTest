package main

import "publicChainTest/client"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 8:47
 **/

func main() {
	//命令行交互入口
	var cli client.Cli
	//执行  --> 命令行输入命令
	cli.Run()
}
