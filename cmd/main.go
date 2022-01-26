package main

import (
    "github.com/gin-gonic/gin"
	"BookStore/controller"
)

func main(){
    server:=gin.Default()
    server.POST("/booksData",controller.GetMaxWordUsed)
    server.Run(":8080")
}