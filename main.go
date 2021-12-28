package main

import (
	sever "YU/Han"
	mysql1 "YU/mysql"
	"fmt"
	"time"
)

func main()  {

	err :=sever.Start()
	if err!= nil{
		panic(err)
	}
	TimeTorefrush()

}
//创建刷新每日数据定时任务
func TimeTorefrush() {
	for {
		now := time.Now()  //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))//计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		fmt.Println("凌晨刷新实现在:",time.Now())
		//以下为定时执行的操作
		mysql1.Delecttodaydate()
	}
}













