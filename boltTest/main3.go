package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"math"
	"os"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/14 18:01
 * @Description:
	项目涉及用户：linfeifei  xiaojie  所有的交易都是围绕这两个用户
	主要功能：
		1、登录
		2、转账（存款，取款）
		3、查询余额
		4、帮助说明文档
	注意事项：
		这里初始化的时候是两个用户分别分配了1000元，限定死了，
		但是我们的数据存储是boltDB持久化的，
		所以，在后期我们无论操作多少次程序，我们只会从持久化文件中读取数据，
		然后进行金额的修改，从而实现交易和简易ATM
 **/

const USERNAME = "linfeifei"
const USERNAME2 = "xiaojie"
const PASSWORD = 123456
const MONEY = 1000.00               //初始余额
const BALANCE_PATH = "./balance.db" //余额文件
const BANLANCE_BUCKET = "balance"   //key  --  value余额

var Aban float64
var Bban float64
var db *bolt.DB

func main() {
	fmt.Println("welcome to MyACM.....")
	Run()
	defer db.Close()
}

func Run() {
	//初始化
	createMoney()

	//判断输入的长度
	if len(os.Args) < 2 {
		return
	}

	//main.exe  ...
	switch os.Args[1] {
	case "login":
		login()
	case "transfer":
		transfer()
	case "save":
		save()
	case "draw":
		draw()
	case "balance":
		balance()
	case "help":
		help()
	case "exit":
		return
	default:
		fmt.Println("功能还在开发中...")
	}

}

/*
	初识状态：两个账户都有1000元，并存入数据库bolt
*/
func createMoney() (err error) {
	db, err = bolt.Open(BALANCE_PATH, 0600, nil)
	if err != nil {
		fmt.Println("打开数据库失败", err.Error())
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BANLANCE_BUCKET))
		if bucket == nil {
			bk, err2 := tx.CreateBucket([]byte(BANLANCE_BUCKET))
			if err2 != nil {
				fmt.Println(err.Error())
				return err2
			}
			//添加数据  linfeifei:1000   xiaojie:1000
			m1 := Float64ToByte(MONEY)
			bk.Put([]byte(USERNAME), m1)
			bk.Put([]byte(USERNAME2), m1)

		} else {
			//直接使用
			bucket := tx.Bucket([]byte(BANLANCE_BUCKET))
			//获取余额
			get := bucket.Get([]byte(USERNAME))
			get2 := bucket.Get([]byte(USERNAME2))
			//赋值给全局变量
			Aban = ByteToFloat64(get)
			Bban = ByteToFloat64(get2)
			/*fmt.Println(Aban)
			fmt.Println(Bban)*/
		}
		return nil
	})
	return err
}

/*
	登录
*/
func login() {
	//login --username "linfeifei" --password 123456
	login := flag.NewFlagSet("login", flag.ExitOnError)
	//获取参数
	username := login.String("username", "", "用户名")
	password := login.Int("password", 0, "密码")
	//解析
	login.Parse(os.Args[2:])

	//判断是否一致
	if *username == USERNAME && *password == PASSWORD {
		fmt.Println("登录成功")
	} else {
		fmt.Println("用户名密码错误，请重试!")
		return
	}
}

/*
	转账
*/
func transfer() error {
	transfer := flag.NewFlagSet("transfer", flag.ExitOnError)
	money := transfer.Float64("money", 0.0, "转账金额")
	//who := transfer.String("to", "", "转账给谁")
	//解析
	transfer.Parse(os.Args[2:])
	fmt.Println(*money)

	/*查询数据库，看余额够不够*/
	err := db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BANLANCE_BUCKET))
		if bk == nil {
			return errors.New("余额桶为空")
		}
		if Aban < *money {
			fmt.Println("您的账户余额不足，补不能进行转账！")
			return nil
		}
		//本人更新数据
		Aban -= *money
		fmt.Println("付款人余额：", Aban)
		//余额转型存入数据库
		n := Float64ToByte(Aban)
		bk.Put([]byte(USERNAME), n)

		//收款人更新数据
		Bban += *money
		n2 := Float64ToByte(Bban)
		bk.Put([]byte(USERNAME2), n2)
		fmt.Println("收款人余额：", ByteToFloat64(bk.Get([]byte(USERNAME2))))
		return nil
	})
	return err
}

/*
	存钱 (基于linfeifei)
*/
func save() error {
	save := flag.NewFlagSet("save", flag.ExitOnError)
	m := save.Float64("money", 0.0, "存款金额")
	save.Parse(os.Args[2:])
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BANLANCE_BUCKET))
		if bucket == nil {
			return errors.New("余额桶为空")
		}
		Aban += *m
		n := Float64ToByte(Aban)
		//更新余额
		bucket.Put([]byte(USERNAME), n)
		return nil
	})
	return err
}

/*
	取钱 (基于linfeifei)
*/
func draw() error {
	draw := flag.NewFlagSet("draw", flag.ExitOnError)
	m := draw.Float64("money", 0.0, "取款金额")
	draw.Parse(os.Args[2:])
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BANLANCE_BUCKET))
		if bucket == nil {
			return errors.New("余额桶为空")
		}
		Aban -= *m
		n := Float64ToByte(Aban)
		//更新余额
		bucket.Put([]byte(USERNAME), n)
		return nil
	})
	return err

}

/*
	查询余额
*/
func balance() (float64, error) {
	ban := 0.0
	err := db.Update(func(tx *bolt.Tx) error {
		bk1 := tx.Bucket([]byte(BANLANCE_BUCKET))
		if bk1 == nil {
			return errors.New("余额桶为空")
		}
		balance := bk1.Get([]byte(USERNAME))
		ban = ByteToFloat64(balance)
		return nil
	})
	fmt.Println("本人余额为：", ban)
	return ban, err
}

/*
	查看帮助菜单
*/
func help() {
	menu()
}
/*
	帮助文档
*/
func menu() {
	fmt.Println("**********功能示例*********")
	fmt.Println("登录: login --username 'XXXX' --password 'xxxxx' ")
	fmt.Println("转账：transfer --money xxx")
	fmt.Println("存款：save --money xxx")
	fmt.Println("取款：draw --money xxx")
	fmt.Println("查看余额：balance")
	fmt.Println("帮助菜单：help")
	fmt.Println("退出:exit")
}

/*
	float 转为 []byte
*/
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

/*
	byte 转为 float64
*/
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
