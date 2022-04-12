package entity

import (
	"bytes"
	"encoding/gob"
	"time"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 8:48
 **/
type Block struct {
	//交易信息
	Data []byte
	//上一个区块的hash
	PrevHash []byte
	//时间戳
	TimeStamp int64
	//随机数
	Nonce int64
	//当前区块hash
	NowHash []byte
}

/*
	创建区块
*/
func NewBlock(data []byte, prevHash []byte) *Block {
	block := Block{
		Data:      data,
		PrevHash:  prevHash,
		TimeStamp: time.Now().Unix(),
	}

	//利用pow工人，计算出随机数
	pow := NewPow(block.Data, block.PrevHash, block.TimeStamp)

	//返回hahs值，算出随机数
	hahs, nonce := pow.Run()

	//赋值给block
	block.NowHash = hahs
	block.Nonce = nonce
	//返回区块
	return &block
}

/*
	序列化：将结构体数据 ---> 字节[]byte切片数据
*/
func (block *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer //写，输出流
	//创建序列化对象,并声明接受者
	en := gob.NewEncoder(&result)
	/*
		json.Marshal()  //序列化
		json.Unmarshal() //反序列化
	*/
	//序列化
	err := en.Encode(block)
	if err != nil {
		return nil, err
	}
	//返回序列化后的数据
	return result.Bytes(), nil
}

/*
	反序列化：将字节[]byte切片数据  ---> 转为结构体数据
*/
func (b *Block) DeSerialize(data []byte) (*Block, error) {
	var reader = bytes.NewReader(data) //读，输入流
	//创建反序列化对象，
	de := gob.NewDecoder(reader)
	//声明接收的结构体对象
	var block *Block
	err := de.Decode(&block)
	if err != nil {
		return nil, err
	}
	return  block,nil
}

