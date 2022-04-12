package tools

import "os"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 14:33
 **/
/*
	判断文件是否存在
*/
func FileExits( path string)bool{
	/*
		返回值1：文件描述
		返回值2：文件存在err为空，不存在err为错误
	*/
	_, err := os.Lstat(path)
	/*
		os.IsNotExist
			错误存在 --->返回true，代表文件不存在
			错误不存在 --->  返回false，代表文件存在
	*/
	return !os.IsNotExist(err)
}