package handler

import (
	"github.com/FlyingDuck/library/dal"
	"github.com/FlyingDuck/library/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Content-Type: application/x-www-form-urlencoded
func AddHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultPostForm("id", "0"), 10, 64)
	bookname := c.PostForm("book_name")
	author := c.PostForm("author")
	category := c.PostForm("category")
	location := c.PostForm("location")
	notice := c.PostForm("notice")
	source := c.PostForm("source")
	state := c.PostForm("state")
	keywordArrayStr := c.PostForm("keyword_list")
	imageArrayStr := c.PostForm("image_list")
	abstract := c.PostForm("abstract")
	isbn := c.PostForm("isbn")

	keywords := strings.Split(keywordArrayStr, ",")
	images := strings.Split(imageArrayStr, ",")

	book := &model.Book{
		ID:       id,
		BookName: bookname,
		Author:   author,
		Category: category,
		Location: location,
		Notice:   notice,
		Source:   source,
		State:    state,
		Keywords: keywords,
		Images:   images,
		Abstract: abstract,
		ISBN:     isbn,
	}

	bookDal := dal.NewBookDal()
	if book.ID > 0 {
		dalErr := bookDal.Update(book)
		if dalErr != nil {
			log.Fatalln(dalErr)
		} else {
			log.Println("update book")
		}
	} else {
		ID, dalErr := bookDal.Add(book)
		if dalErr != nil {
			log.Fatalln(dalErr)
		} else {
			log.Printf("new book(ID=%d) info\n", ID)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"bookID": book.ID,
	})
}
