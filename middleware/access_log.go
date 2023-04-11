package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatalf("read body failed at Before,err:%s", err.Error())
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		log.Printf("request: url=%s, body=%s\n", c.Request.URL, string(body))

		// response
		blw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		c.Next()
		log.Printf("response: status=%d, body=%s\n", c.Writer.Status(), blw.body.String())
	}
}
