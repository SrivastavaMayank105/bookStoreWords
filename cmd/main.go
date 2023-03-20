package main

import (
    "github.com/gin-gonic/gin"
	"BookStore/controller"
	"fmt"
)

func main(){
    fmt.Println("hello")
    server:=gin.Default()
    server.POST("/booksData",controller.GetMaxWordUsed)
    server.Run(":8080")
}