package entity

import (
	"bytes"
	"math/big"
	"publicChainTest/tools"
	"strconv"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 9:32
 **/
/*
	工作量证明pow，也就是一个工人，计算随机数的
*/
//难度系数
const BITS = 10

type ProofOfWork struct {
	//给那个区块工作，那就需要区块里面的信息
	Data      []byte
	PrevHash  []byte
	TimeStamp int64

	TargetHash *big.Int //系统给定的目标hash
}

/*
	创建工作量证明pow对象，并返回，并给target赋值
	1、实例化一个结构体
	2、算出target，并赋值
	3、返回pow结构体
*/
func NewPow(data []byte, prevHash []byte, timeStamp int64) *ProofOfWork {
	//系统目标值
	target := big.NewInt(1)
	//目标值左移，20个0 + 1 + 256-Bit-1个0
	target =target.Lsh(target,256-BITS-1)//移动的数字，移动的位数
	pow := ProofOfWork{
		Data: data,
		PrevHash: prevHash,
		TimeStamp: timeStamp,
		TargetHash: target, //系统目标hash
	}
	return &pow
}

/*
	计算随机数，并把区块hash值 和 随机数 返回
*/
func (pow *ProofOfWork)Run()([]byte,int64) {
	//随机数
	var nonce int64
	nonce = 0
	 //从0开始找
	//时间戳转为[]byte类型
	time := []byte(strconv.FormatInt(pow.TimeStamp,10))
	/*
		循环比对中做的事情：
			1.  准备数据
			2.  用 SHA-256 对数据进行哈希
			3.  将哈希转换成一个大整数
		 	4.将这个大整数与目标进行比较
	*/
	for{
		//随机数也转为[]byte类型
		nonceByte := []byte(strconv.FormatInt(nonce,10))

		//拼接：交易信息 + 上一个区块hash + 时间戳 + 随机数
		byteInfo := bytes.Join([][]byte{pow.Data, pow.PrevHash, time, nonceByte}, []byte{})

		//计算成hash值
		hash := tools.GetSHA256Hash(byteInfo)

		//把hash转为bigInt
		num := big.NewInt(0)
		num = num.SetBytes(hash)
		/*
			规则：
				if(a < target){
				//a：区块的hash值   target：系统给定的hahs值
				}
		*/
		if num.Cmp(pow.TargetHash) == -1 {//返回-1
			//找到了，返回hash值和随机数
			return hash,nonce
		}
		//每算错一次，随机数++
		nonce++
	}
	return nil,0
}
