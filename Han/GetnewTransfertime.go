package Hander

import (
	geth "YU/connectgeth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	time2 "time"
)

func GetnewTransfertime( c *gin.Context) {
	//_,add:=geth.Getaccout()
	add:=common.HexToAddress(c.PostForm("account"))
	ins :=geth.Getsmartcontract()
	header,err :=geth.GetBlockNumber()
	if err !=nil {
		panic(err)
	}
	transfertime :=geth.Gettransfertime(ins,add,header)
	time := time2.Unix(transfertime.Int64(),0)
	fmt.Println("最新交易时间为",time.Format("2006-01-02 15:04:05"))
    tohtml(c,time)
}
