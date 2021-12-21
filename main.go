package main

import (
	sever "YU/Han"
)

func main()  {

	//_,add:=geth.Getaccout()
	//ins :=geth.Getsmartcontract()
	//header,err :=geth.GetBlockNumber()
	//if err != nil{
	//	panic(err)
	//}
	//
	//Txopts :=geth.GetTxopts()
	////fmt.Println("============",Txopts)
	//transfermessage := geth.Dotransfer(ins,add,Txopts)
	//fmt.Println("======================")
	//fmt.Println(transfermessage)
	// fmt.Println("=====================")
	//balance :=geth.GetBalanceUser(ins,add,header)
	//fmt.Println("账户余额为",balance)
	//
	//balancecontract :=geth.GetBalancecontract(ins,add,header)
	//fmt.Println("水龙头余额",balancecontract)
	//
	//transfertime :=geth.Gettransfertime(ins,add,header)
	//time := time2.Unix(transfertime.Int64(),0)
	//fmt.Println("最新交易时间为",time.Format("2006-01-02 15:04:05"))
	//
	//
	// transfercount:=geth.Gettransfercount(ins,add,header)
	// fmt.Println("用户账户得币总数量",transfercount)
	//
	//todaytoken:=geth.Gettodaytoken(ins,add,header)
	//fmt.Println("今日用户所获取的币数",todaytoken)

	err :=sever.Start()
	if err!= nil{
		panic(err)
	}
}












