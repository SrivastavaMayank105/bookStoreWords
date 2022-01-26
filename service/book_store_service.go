package service

import (
	"github.com/gin-gonic/gin"	
	"BookStore/mapper"
	"errors"
	"strings"
	"sort"
)

type BookStoreWords interface{
	GetMaxWordUsed(ctx *gin.Context,reqest mapper.BookRequestParams) ([]string,error)
}
type bookStoreMaxWords struct{
	bookWords string
	wordsCount int
}

func NewBookStoreWordsService()BookStoreWords{
	return bookStoreMaxWords{}
}

func (bkmw bookStoreMaxWords)GetMaxWordUsed(ctx *gin.Context,request mapper.BookRequestParams)([]string,error){
	removeTrim:=strings.TrimSpace(request.BookText)
	if len(removeTrim)==0{
		return nil,errors.New("invalid input")
	}

	requestTrimSpace:=strings.Split(request.BookText," ")
	response:=make(map[string]int)
	for i :=range requestTrimSpace{
		value,ok:=response[requestTrimSpace[i]];if !ok{
			response[requestTrimSpace[i]]=1
		}else{
			response[requestTrimSpace[i]]=value+1
		}
	}
	maxFrequenceWords:=bkmw.filterTheMostFrequenceWords(response)
	return maxFrequenceWords,nil
}

func (bkmw bookStoreMaxWords)filterTheMostFrequenceWords(res map[string]int)[]string{
	countWords:=make([]bookStoreMaxWords,0,len(res))
	for value,ok:=range res{
		countWords=append(countWords,bookStoreMaxWords{bookWords:value,wordsCount:ok})
	}

	sort.SliceStable(countWords, func(i, j int) bool {
		return countWords[i].wordsCount > countWords[j].wordsCount
	})

	maxTenWords:=make([]string,0,10)
	for i := 0; i < len(countWords) && i < 10; i++ {
		maxTenWords=append(maxTenWords,countWords[i].bookWords)
	}
	return maxTenWords
}