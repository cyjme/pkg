package pkg

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBindWith(request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "parse Request Error",
			"error": err.Error(),
		})
		log.Println("ParseRequest Result", request)
		log.Println("ParseRequest Error", err.Error())
		return err
	}
	return nil
}

func SuccessResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}
