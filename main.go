package main

import (
	"contract/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	bootstrap.SetupRoute(r)
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
