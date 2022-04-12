package entity

import (
	"errors"
	"github.com/boltdb/bolt"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 8:48
 **/
type BlockChain struct {
	DB       *bolt.DB
	LastHash []byte
}

const BLOCKCHAIN_DB_PATH = "./blockChain.db"  //区块链数据库文件路径
const BUCKET_BLOCK =  "chain_blocks"  //存区块的桶名字
const  BUCKET_LAST =  "chain_last"
const  LAST_HASH = "last_block"
/*
	创建区块链 ---> 改成blot区块链数据库
		1、打开数据库
		2、Update存入数据
		3、先直接使用桶，判断桶是否存在，不存在创建，(避免出现桶存在报错问题)
		4、桶不存在: 1.创建桶。 2.获取到创世区块，存入桶中。  3.创建桶2,存入最后一个区块的hash。
		5、桶存在: 直接使用那个桶2，获取到最后一个区块的hash
		6、给区块链赋值: db对象 + 最后一个区块hash
*/
func NewBlockChain(data []byte) (*BlockChain, error) {
	var lastHash []byte
	//打开数据库
	blotDB, err := bolt.Open(BLOCKCHAIN_DB_PATH, 0600, nil)
	if err != nil {
		return nil, err
	}
	//存数据到区块链中
	blotDB.Update(func(tx *bolt.Tx) error {
		//先试用桶，没有桶再创建桶
		bucket := tx.Bucket([]byte(BUCKET_BLOCK))
		if bucket == nil { ////如果桶为空，说明还没有区块链，就要创建区块链  桶1 = 区块链
			//创建桶1，存区块链的
			bk, err := tx.CreateBucket([]byte(BUCKET_BLOCK))
			if err!= nil {
				return err
			}
			//先获取到创世区块
			genesis := NewGenesisBlock(data)
			//创世区块序列化成[]byte字节数据
			genesisByte, err := genesis.Serialize()
			if err != nil {
				return err
			}
			//添加到桶中  key:哈希  value:区块
			bk.Put(genesis.NowHash, genesisByte)

			//第二个桶2，存储最后一个区块的hash值
			bk2, err := tx.CreateBucket([]byte(BUCKET_LAST))
			if err != nil {
				return err
			}
			//桶2存入数据
			bk2.Put([]byte(LAST_HASH),genesis.NowHash)
			//更新区块链的lastHash
			lastHash = genesis.NowHash

		}else { //有桶了，直接使用桶2，获取到最后一个区块的hash
			bk2 := tx.Bucket([]byte(BUCKET_LAST))
			//获取到最后一个hash值
			lastHash = bk2.Get([]byte(LAST_HASH))
		}
		return nil
	})
	//以上都是准备工作，这里是给区块链结构赋值，也就是创建区块链
	blockChain := 	BlockChain{
		DB: blotDB,
		LastHash: lastHash,
	}
	return &blockChain,nil
}

/*
	创世区块
*/
func NewGenesisBlock(data []byte) *Block {
	//创世区块 (交易信息data,上一个区块hash:32个0特殊化)
	return NewBlock(data,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}

/*
	添加区块到区块链
*/
func (bc *BlockChain)AddBlockToChain(data []byte) error{
	//创建区块
	block := NewBlock(data,bc.LastHash)
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil {
			return errors.New("还没创建桶呢！")
		}
		blockByte,err := block.Serialize()
		if err != nil {
			return err
		}
		//不为空直接用
		bk.Put(block.NowHash,blockByte)

		//更新桶2
		bk2 := tx.Bucket([]byte(BUCKET_LAST))
		if bk2 == nil {
			return  errors.New("还没创建桶2呢！")
		}
		bk2.Put([]byte(LAST_HASH),block.NowHash)
		//更新区块链
		bc.LastHash = block.NowHash
		return nil
	})
	return err
}
/*
	创建迭代器，返回迭代器
*/
func (bc *BlockChain)Iterator()*ChainIterator{
	chainInteractor := ChainIterator{
		DB: bc.DB,
		currentHash: bc.LastHash,
	}
	return &chainInteractor
}
/*
	获取所有的区块对象
*/
func (bc *BlockChain)GetAllBlocks()(blocks []*Block,err error){
	iterator := bc.Iterator()
	for  {
		if iterator.HasNext() {
			//反序列化添加到区块切片中
			next, err := iterator.Next()
			if err != nil {
				return nil,err
			}
			blocks = append(blocks,next )
		}else {
			break
		}
	}
	return blocks,nil
}
