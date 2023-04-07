package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-demo/component"
	"log"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// print err msg
			log.Printf("panic: %v\n", r)
			// debug.PrintStack()
			// response same struct
			c.JSON(400, component.RestResponse{Code: -1, Message: fmt.Sprintf("%v", r)})
		}
	}()

	c.Next()
}
