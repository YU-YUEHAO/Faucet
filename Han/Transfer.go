package Hander

import (
	geth "YU/connectgeth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	a "YU/mysql"
)
func Trasfer(c *gin.Context) {
	add :=common.HexToAddress(c.PostForm("account"))
	ins :=geth.Getsmartcontract()
	Txopts :=geth.GetTxopts()
    transfermessage := geth.Dotransfer(ins,add,Txopts)
    fmt.Println(transfermessage)
    a.Insert(add,1)
    tohtml(c,transfermessage)
    fmt.Println("完成")
}
