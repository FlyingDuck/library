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

	bookDal := dal.NewBookDal()
	dalErr := bookDal.DelByID(bookID)
	if dalErr != nil {
		log.Fatalln(dalErr)
	}

	c.JSON(http.StatusOK, gin.H{})
}
