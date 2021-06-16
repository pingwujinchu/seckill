package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Health ...
func Health(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{"code": "0", "msg": "success"})
}

//Version ...
func Version(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{"code": "0", "msg": "success", "version": "0.1"})
}
