package client

import (
	"flag"
	"fmt"
	"os"
	"publicChainTest/entity"
	"publicChainTest/tools"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/12 8:56
 **/
/*
	用户的交互入口：
		只用于负责读取用户传递的命令和参数
		并进行解析
		传递解析参数，调用对应的功能
*/
type Cli struct {
	//首先需要有区块链
	bc *entity.BlockChain
}

func (cl *Cli) Run() {
	//使用区块链
	chain, _ := entity.NewBlockChain([]byte("创世区块"))
	//延迟关闭
	defer chain.DB.Close()
	cl.bc = chain
	if cl.bc == nil {
		fmt.Println("区块链db不存在")
	}
	//判断输入的长度
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "createblockchain":
		cl.createChain()
	case "addblock":
		cl.addBlock()
	case "printchain":
		cl.printChain()
	case "getblockcount":
		cl.getBlockCount()
	case "getblockinfo":
		cl.getBlockInfo()
	case "getlastblock":
		cl.getLastBlock()
	case "getfirstblock":
		cl.getFirstBlock()
	case "getallblock":
		cl.getAllBlock()
	case "help":
		cl.help()
	default:
		fmt.Println("please check it！not have this function~")
		//退出
		os.Exit(1)
	}

}

//对应上面的列表
/*
	创建区块链
*/
func (cl *Cli) createChain() {
	//声明成功能
	cre := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	//获取参数
	data := cre.String("data", "", "创建区块链的交易信息")
	//解析
	cre.Parse(os.Args[2:])

	//判断区块链是否存在
	exits := tools.FileExits("blockChain.db")
	if exits {
		fmt.Println("区块链已经存在，不能再创建了！")
		return
	}
	_, err := entity.NewBlockChain([]byte(*data))
	if err != nil {
		fmt.Println("区块链创建失败")
		return
	}
	fmt.Println("区块链创建成功")
}

/*
	添加区块
*/
func (cl *Cli) addBlock() {
	//判断区块链是否存在
	exits := tools.FileExits("blockChain.db")
	if !exits {
		fmt.Println("区块链不存在，请创建区块链后添加区块！")
		return
	}
	add := flag.NewFlagSet("addblock", flag.ExitOnError)
	data := add.String("data", "", "添加区块的交易信息")
	add.Parse(os.Args[2:])
	err := cl.bc.AddBlockToChain([]byte(*data))
	if err != nil {
		fmt.Println("添加区块失败")
		return
	}
	fmt.Println("添加区块成功")
}

/*
	迭代区块链
*/
func (cl *Cli) printChain() {
	//判断区块链是否存在
	exits := tools.FileExits("blockChain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlocks()
	if err != nil {
		fmt.Println("获取区块链对象失败！",err.Error())
		return
	}
	for _, block := range blocks {
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("PrevHash:%x\n",block.PrevHash)
		fmt.Printf("Hash:%x\n",block.NowHash)
	}
	fmt.Println("遍历完成！！")
}

/*
	获取区块数量
*/
func (cl *Cli) getBlockCount() {
	exits := tools.FileExits("blockChain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlocks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("区块总数量：",len(blocks))
}

/*
	获取指定区块的信息
*/
func (cl *Cli) getBlockInfo() {
	/*
		问题：我们的blotDB只能桶key获取到值，能不能根据具体的vlaue获取到整个信息呢？
	*/
}

/*
	获取第一个区块
*/
func (cl *Cli) getFirstBlock() {
	exits := tools.FileExits("blockChain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlocks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Data:%s\n",blocks[len(blocks)-1].Data)
	fmt.Printf("PrevHash:%x\n",blocks[len(blocks)-1].PrevHash)
	fmt.Printf("Hash:%x\n",blocks[len(blocks)-1].NowHash)
}
/*
	获取最后一个区块
*/
func (cl *Cli) getLastBlock() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlocks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//最后一个切片对象也就是最后一个区块
	fmt.Printf("Prev. hash: %x\n", blocks[0].PrevHash)
	fmt.Printf("Data: %s\n", blocks[0].Data)
	fmt.Printf("Hash: %x\n", blocks[0].NowHash)
}
/*
	获取所有的区块
*/
func (cl *Cli) getAllBlock() {
	exits := tools.FileExits("blockChain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlocks()
	if err != nil {
		fmt.Println("所有区块获取失败！",err.Error())
		return
	}
	fmt.Println("获取到的区块对象个数：",len(blocks))
}

/*
	帮助菜单
*/
func (cl *Cli) help() {
	fmt.Println("main.exe Command --data ?")
	fmt.Println("Has the following Command:")
	fmt.Println("\t \t createBlockchain --data Transaction information of Genesis block")
	fmt.Println("\t \t addblock --data Transaction information of this block")
	fmt.Println("\t \t getblockinfo --hash The hash of this block")
	fmt.Println("\t \t printchain")
	fmt.Println("\t \t getblockconut")
	fmt.Println("\t \t getlastblock")
	fmt.Println("\t \t getfirstblock")
	fmt.Println()
}
