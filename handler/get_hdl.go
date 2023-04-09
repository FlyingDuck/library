package handler

import (
	"github.com/FlyingDuck/library/dal"
	"github.com/FlyingDuck/library/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetHandler(c *gin.Context) {
	bookID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	bookDal := dal.NewBookDal()
	book, dalErr := bookDal.FindByID(bookID)
	if dalErr != nil {
		log.Fatalln(dalErr)
	} else {
		if book != nil {
			book.Category = model.CategoryMapping[book.Category]
			book.Location = model.LocationMapping[book.Location]
			book.Source = model.SourceMapping[book.Source]
			book.State = model.StateMapping[book.State]
		}
	}

	c.JSON(http.StatusOK, book)
}
