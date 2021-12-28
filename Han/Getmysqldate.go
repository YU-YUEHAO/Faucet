package Hander

import (
	 m "YU/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Todaydate(c *gin.Context) {
     todaydate :=  m.Selecttodaydate()
	 tohtml(c,todaydate)
}
func Alldate(c *gin.Context)  {

  alldate := m.Selectalldate()
  fmt.Println("allll",alldate)
  tohtml(c,alldate)
}
