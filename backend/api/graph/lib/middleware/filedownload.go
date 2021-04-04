package middleware

import (
	"fmt"
	"net/http"

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
	fpath := fmt.Sprintf("%s/%s", userIDDirectory, filename)
	fmt.Printf("%s\n", fpath)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./documents/" + fpath)
}
