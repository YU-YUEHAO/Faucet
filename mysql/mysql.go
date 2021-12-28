package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


var db *sql.DB

func init()  {
	var err error
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db,err=sql.Open("mysql","root:121@tcp(127.0.0.1:3306)/gomysql") // 设置连接数据库的参数
	if err!= nil{
		panic(err)
	}
	//设置数据库最大连接数
	//db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	//db.SetMaxIdleConns(10)
	//defer db.Close()                                                   //关闭数据库
	err = db.Ping()                                                    //连接数据库
	if err!=nil {
		fmt.Println("数据库连接失败")
	}
}
//插入所有数据
func Insert(add string,number int)  error {
	var err error
	var sqlstate2 = `SELECT MAX(id) FROM alldate`
	rows,err :=db.Query(sqlstate2)
	var maxid int
	if err!=nil {
		panic(err)

	}
	for rows.Next() {
		err=rows.Scan(&maxid)

	}
	fmt.Println(maxid)
	var sqlStr = `INSERT INTO alldate (id,address,transationtime,transationNumber) VALUES (?,?,?,?)`
	_, err = db.Exec(sqlStr,maxid+1,add,time.Now().Format("2006-01-02 15:04:05"),number)
	if err != nil {
		return err
	}
 fmt.Println("总数据插入成功")
	return err
}
//插入当天数据
func Inserttoday(add string,number int)  error{
	var err error
	var sqlstate2 = `SELECT MAX(id) FROM todaydate`
	rows,err :=db.Query(sqlstate2)
	var maxid int
	if err!=nil {
		panic(err)

	}
	for rows.Next() {
		err=rows.Scan(&maxid)

	}
	var sqlStr = ` INSERT INTO todaydate (id,Transation,toAddress,Transationtime) VALUES (?,?,?,?) `
	_,err =db.Exec(sqlStr,maxid+1,number,add,time.Now().Format("2006-01-02 15:04:05"))
	if err!=nil{
		panic(err)
	}
	fmt.Println("今天数据插入成功")
   return err
}
//查询当天数据
func Selecttodaydate() int {
	var sqlStr="SELECT COUNT(Transation) FROM todaydate"
	rows,err :=db.Query(sqlStr)
	fmt.Println("rows=====>",rows)
	if err!=nil {
		panic(err)
	}
		 var date int
	     for rows.Next() {
		 err :=rows.Scan(&date)
		if err!=nil {
			 panic(err)
		}
		fmt.Println(date)
	}
	 return date

}
//查询所有数据
func Selectalldate() int {
	var sqlStr=` SELECT COUNT(transationNumber) FROM alldate`
	//alldate,err :=db.Prepare(sqlStr)
	rows,err :=db.Query(sqlStr)
	if err!=nil {
		panic(err)
	}
	var date int
	for rows.Next() {
		err :=rows.Scan(&date)
		if err!=nil {
			panic(err)
		}
		fmt.Println(date)
	}
	return date
}

//每日刷新当天数据
func Delecttodaydate()  error{
	var sqlStr=`DELETE FROM todaydate`
	_,err:=db.Exec(sqlStr)
	if err!=nil {
		panic(err)
	}
	return err
}

