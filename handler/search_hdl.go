package handler

import (
	"github.com/FlyingDuck/library/dal"
	"github.com/FlyingDuck/library/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SearchHandler(c *gin.Context) {
	pageNum, _ := strconv.ParseInt(c.DefaultQuery("pn", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("ps", "10"), 10, 64)

	shelf := &BookShelf{}

	bookDal := dal.NewBookDal()
	books, total, dalErr := bookDal.FindByModify(int(pageSize*(pageNum-1)), int(pageSize))
	if dalErr != nil {
		log.Fatalln(dalErr)
	} else {
		for _, book := range books {
			book.Category = model.CategoryMapping[book.Category]
			book.Location = model.LocationMapping[book.Location]
			book.Source = model.SourceMapping[book.Source]
			book.State = model.StateMapping[book.State]
		}
		shelf.Books = books
		shelf.Total = total
		shelf.PageSize = pageSize
		shelf.PageNum = pageNum
	}

	c.JSON(http.StatusOK, shelf)

}

type BookShelf struct {
	Total    int64         `json:"total"`
	Books    []*model.Book `json:"books"`
	PageNum  int64         `json:"page_num"`
	PageSize int64         `json:"page_size"`
}
