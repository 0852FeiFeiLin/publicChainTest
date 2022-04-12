package tools

import "crypto/sha256"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 10:01
 **/
func GetSHA256Hash(data []byte)([]byte){
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
