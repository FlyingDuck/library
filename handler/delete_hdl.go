package handler

import (
	"github.com/FlyingDuck/library/dal"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteHandler(c *gin.Context) {
	bookID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	bookIDArray := make([]int64, 0)
	bindErr := c.ShouldBindJSON(&bookIDArray)
	if bindErr != nil {
		log.Println(bindErr)
	}
	bookIDArray = append(bookIDArray, bookID)

	bookDal := dal.NewBookDal()
	dalErr := bookDal.DelByIDs(bookIDArray)
	if dalErr != nil {
		log.Println(dalErr)
	}

	c.JSON(http.StatusOK, gin.H{})
}
