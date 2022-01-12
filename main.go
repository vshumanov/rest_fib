package main

import (
	"math/big"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	engine := gin.Default()
	engine.GET("/fibbonaci", handleGet)
	return engine
}

func handleGet(c *gin.Context) {
	if req_n, ok := c.GetQuery("n"); ok {
		n := new(big.Int)
		n, ok := n.SetString(req_n, 10)
		if !ok || n.Sign() == -1 {
			c.JSON(400, gin.H{
				"error": "Can't compute fibbonaci with the provided parameter",
			})
			return
		}

		c.JSON(200, gin.H{
			"result": fibbonaci(n),
		})
	}
}
