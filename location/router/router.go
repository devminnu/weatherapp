package router

import (
	"time"

	ginzap "github.com/akath19/gin-zap"
	"github.com/devminnu/weatherapp/location/logger"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init() {
	r = gin.New()
	//Zap logger
	logger := logger.Get()
	//Add middleware to Gin, requires sync duration & zap pointer
	r.Use(ginzap.Logger(3*time.Second, logger))
	//Other gin configs
	r.Use(gin.Recovery())
}

func Get() *gin.Engine {
	return r
}
