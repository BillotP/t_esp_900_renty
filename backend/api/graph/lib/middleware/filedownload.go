package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FileDownload(c *gin.Context) {
	filename, ok := c.Params.Get("filename")
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
	}
	userIDDirectory, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
	}
	scopeDirectory, ok := c.Params.Get("scope")
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
	}
	fpath := fmt.Sprintf("%s/%s/%s", scopeDirectory, userIDDirectory, filename)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	dir, _ := os.Getwd()
	fmt.Println(dir + "/data/" + fpath)
	c.File(dir + "/data/" + fpath)
}
