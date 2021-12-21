package mysql

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"
)


var db *sql.DB

func Init(file string) (err error) {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db,_:=sql.Open("mysql","root:121@(127.0.0.1:3306)/gomysql") // 设置连接数据库的参数
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	defer db.Close()                                                   //关闭数据库
	err = db.Ping()                                                    //连接数据库
	if err!=nil {
		fmt.Println("数据库连接失败")
	}
	return
}
//插入数据
func Insert(add common.Address,number int) (err error) {
	var sqlStr = `INSERT INTO alldate (address,transationtime,transationNumber) VALUES (?,?,?)`
	r, err := db.Exec(sqlStr, add,time.Now().Format("2006-01-02 15:04:05"),number)
	if err != nil {
		return err
	}
 fmt.Println("reer",r)
	return
}
