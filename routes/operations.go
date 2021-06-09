package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func sumXY(c *gin.Context) {
	x, err1 := strconv.Atoi(c.Query("x"))
	y, err2 := strconv.Atoi(c.Query("y"))
	if err1 != nil {
		fmt.Println("err1", err1)
		c.JSON(http.StatusOK, err1.Error())
		return
	}
	if err2 != nil {
		fmt.Println("err1", err2)
		c.JSON(http.StatusOK, err2.Error())
		return
	}
	c.JSON(http.StatusOK, x+y)
}


