package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/dongsj/Downloads/骑行.jpeg" \
  -H "Content-Type: multipart/form-data"
*/
func UploadHandler(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./assets/upload/" + file.Filename
	// 上传文件至指定的完整文件路径
	_ = c.SaveUploadedFile(file, dst)

	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"path":     fmt.Sprintf("/assets/upload/%s", file.Filename),
		"URL":      fmt.Sprintf("http://localhost:8080/assets/upload/%s", file.Filename),
	})
}
