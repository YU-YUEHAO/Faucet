package Hander

import (
	geth "YU/connectgeth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetcontractBalance(c *gin.Context) {
	//_,add:=geth.Getaccout()
	add :=common.HexToAddress((c.PostForm("account")))
	fmt.Println(add)
	ins :=geth.Getsmartcontract()
	header,err :=geth.GetBlockNumber()
	if err!=nil{
		panic(err)
	}
		balancecontract :=geth.GetBalancecontract(ins,add,header)
		fmt.Println(balancecontract)
       tohtml(c,balancecontract)
	}
