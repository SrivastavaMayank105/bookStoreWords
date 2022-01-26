package controller

import (
	"github.com/gin-gonic/gin"
	"BookStore/mapper"
	"BookStore/service"
	"net/http"
	"fmt"
)
func GetMaxWordUsed(ctx *gin.Context){
	var request mapper.BookRequestParams

	err:=ctx.ShouldBindJSON(&request);if err !=nil{
		bindErr :=fmt.Errorf("Bad requedt")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,bindErr.Error())
		return
	}
	bookStoreRes:=service.NewBookStoreWordsService()

	response,responseErr:=bookStoreRes.GetMaxWordUsed(ctx,request)
	if responseErr!=nil{
		err:=fmt.Errorf("unable to count the data")
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable,err.Error())
		return
	}

	ctx.JSON(http.StatusOK,response)
}