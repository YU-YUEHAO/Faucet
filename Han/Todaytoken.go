package Hander

import (
	geth "YU/connectgeth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

)

func Gettodaytoken(c *gin.Context) {
	add :=common.HexToAddress((c.PostForm("account")))
	ins :=geth.Getsmartcontract()
	header,err :=geth.GetBlockNumber()
	if err!=nil {
		 panic(err)
	}
	todaytoken:=geth.Gettodaytoken(ins,add,header)
	fmt.Println(todaytoken)
	tohtml(c,todaytoken)
}