
package Hander

import (
"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()
	//配置静态文件
	r.Static("/", "html")

	//设置路由组
	ygh := r.Group("ygh")
	{
		ygh.POST("/dotransfer",Trasfer)
		ygh.POST("/cAmount", GetcontractBalance)
		ygh.POST("/gettoday",Gettodaytoken )
		ygh.POST("/useralltoke",Getalltoken)
		ygh.POST("/gettime", GetnewTransfertime)
		ygh.POST("/getbalaceUser",GetbalanceUser)
		ygh.POST("/todaydata",Todaydate)
		ygh.POST("/alldata",Alldate)
	}
	err := r.Run("127.0.0.1:1111")
	return err
}

func tohtml(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

func returnError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}