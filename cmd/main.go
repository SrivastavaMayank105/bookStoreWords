package main

import (
    "github.com/gin-gonic/gin"
	"BookStore/controller"
	"fmt"
	"os"
)

func main(){
    fmt.Println("hello")
    server:=gin.Default()
    server.POST("/booksData",controller.GetMaxWordUsed)
    value := os.Getenv("PORT")
    server.Run(":"+value)
}