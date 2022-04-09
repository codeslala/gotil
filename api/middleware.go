package api

import (
	"github.com/codeslala/gotil/errs"
	"github.com/gin-gonic/gin"
)

func UrlEncodedContentTypeCheckMW(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType != "application/x-www-form-urlencoded" {
		Result400(c, errs.InvalidArgument.New("wrong Content-Type"))
		return
	}
	c.Next()
}

func JsonContentTypeCheckMW(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		Result400(c, errs.InvalidArgument.New("wrong Content-Type"))
		return
	}
	c.Next()
}
