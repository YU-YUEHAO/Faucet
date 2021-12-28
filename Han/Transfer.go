package Hander

import (
	geth "YU/connectgeth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	a "YU/mysql"
)
func Trasfer(c *gin.Context)  {
	add :=common.HexToAddress(c.PostForm("account"))
	ins :=geth.Getsmartcontract()
	Txopts :=geth.GetTxopts()
	fmt.Println(add)
    transfermessage, err1 := geth.Dotransfer(ins,add,Txopts)
    fmt.Println(err1)
	if err1 !=nil {
		returnError(c,err1)
		return
	}
    fmt.Println(transfermessage)
    add1:=add.String()
	err := a.Insert(add1, 1)
	if err!=nil {
		panic(err)
	}
	err = a.Inserttoday(add1, 1)
	if err!=nil {
		panic(err)
	}
    tohtml(c,transfermessage)
    fmt.Println("完成")

}
