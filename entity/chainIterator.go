package entity

import (
	"bytes"
	"errors"
	"github.com/boltdb/bolt"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 8:48
 **/
type ChainIterator struct {
	DB          *bolt.DB
	currentHash []byte
}

/*
	获取下一个区块
*/
func (iterator *ChainIterator) Next() (*Block, error) {
	var block *Block
	var err error
	err = iterator.DB.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil {
			return errors.New("还没创建桶呢！")
		}
		//通过currentHash也就是lastHash找到最后一个区块
		BlockBytes := bk.Get(iterator.currentHash)
		//反序列化
		block, err = block.DeSerialize(BlockBytes)
		//更新标志位 = 当前区块的上一个区块，这样下一次迭代出来的就是上一个区块
		iterator.currentHash = block.PrevHash
		return nil
	})
	return block, err
}

/*
	判断是否还有下一个区块(比较标志位 == 创世区块的prevHash)
*/
func (iterator *ChainIterator) HasNext() bool {
	//-1小于   0等于   1大于
 	i := bytes.Compare(iterator.currentHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return i != 0
}
