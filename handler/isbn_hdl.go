package handler

import (
	"fmt"
	"github.com/FlyingDuck/library/model"
	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

/*

http://opac.nlc.cn/F/BFBIIUURCC44ACX7IEB5EVBEQK62BSJIPPAM6EU5XUA4J8DCH8-03132?func=find-b&find_code=ISB&request=9787101052039&local_base=NLC01&filter_code_1=WLN&filter_request_1=&filter_code_2=WYR&filter_request_2=&filter_code_3=WYR&filter_request_3=&filter_code_4=WFM&filter_request_4=&filter_code_5=WSL&filter_request_5=
*/

func ISBNHandler(c *gin.Context) {
	isbn := c.Param("isbn")

	baseURL := "http://opac.nlc.cn/F"

	// 发送Get请求
	//rsp, err := http.Get(baseURL)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//body, err := ioutil.ReadAll(rsp.Body)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//content := string(body)
	//defer rsp.Body.Close()
	resp, err := soup.Get(baseURL)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"message": "base URL request fail"})
		return
	}

	// 解析页面，获取 form 表单中搜索请求完整前缀
	doc := soup.HTMLParse(resp)
	form := doc.Find("form", "name", "form1")
	if form.Pointer == nil {
		c.JSON(http.StatusOK, gin.H{"message": "real base URL not exist"})
		return
	}
	realBaseURL, existing := form.Attrs()["action"]
	if !existing {
		c.JSON(http.StatusOK, gin.H{"message": "real base URL not exist"})
		return
	}

	// 拼接新的数据请求
	dataURL := fmt.Sprintf("%s?func=find-b&find_code=ISB&request=%s&local_base=NLC01&filter_code_1=WLN&filter_request_1=&filter_code_2=WYR&filter_request_2=&filter_code_3=WYR&filter_request_3=&filter_code_4=WFM&filter_request_4=&filter_code_5=WSL&filter_request_5=", realBaseURL, isbn)
	dataResp, err := soup.Get(dataURL)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"message": "data URL request fail"})
		return
	}

	doc = soup.HTMLParse(dataResp)
	div := doc.Find("div", "id", "details2")
	if div.Pointer == nil {
		c.JSON(http.StatusOK, gin.H{"message": "data not exist"})
		return
	}
	trs := div.Find("table").Find("tbody").FindAll("tr")
	book := &model.Book{
		ISBN: isbn,
	}
	for _, tr := range trs {
		tds := tr.FindAll("td")
		if len(tds) != 2 {
			continue
		}
		key := strings.TrimSpace(tds[0].Text())
		value := strings.TrimSpace(tds[1].FullText())
		switch key {
		case "内容提要":
			book.Abstract = value
		case "主题":
			book.Theme = value
		//case "题名":
		//	book.BookName = value
		//case "著者":
		//	book.Author = value
		case "题名与责任":
			parts := strings.Split(value, "/")
			if len(parts) == 2 {
				book.BookName = strings.TrimSpace(parts[0])
				book.Author = strings.TrimSpace(parts[1])
			}
		case "出版项":
			book.Publisher = value
		default:
			continue
		}
	}

	if len(book.BookName) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "data not exist"})
		return
	}

	c.JSON(http.StatusOK, book)
}
