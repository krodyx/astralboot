// Serve web interface

package main

import (
	"github.com/gin-gonic/gin"
)

// WebInterface : provides a web interface for astralboot functions and monitoring
func (wh *WebHandler) WebInterface() {
	wh.router.GET("/", wh.Index)
	//wh.router.GET("/images/:source/:rocket/:imageName", wh.AciImage)
}

func (wh *WebHandler) Index(c *gin.Context) {
	logger.Debug("Index HIT")
}
