package Hander

import (
	geth "YU/connectgeth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)
//用户得到的总币数
func Getalltoken(c *gin.Context) {
	add :=common.HexToAddress(c.PostForm("account"))
	ins :=geth.Getsmartcontract()
	header,err :=geth.GetBlockNumber()
	if err!=nil {
		panic(err)
	}
 transfercount:=geth.Gettransfercount(ins,add,header)
  fmt.Println(transfercount)
    tohtml(c,transfercount)
}
