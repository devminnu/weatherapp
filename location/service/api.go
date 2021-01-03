package service

import (
	"github.com/devminnu/weatherapp/location/router"
)

func Init() {
	r := router.Get()
	r.GET("/getLocationByID", getLocationByID)
}
