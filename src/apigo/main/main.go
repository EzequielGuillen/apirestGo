package main

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"

)

var (

	router = gin.Default()
)

func main() {
	router.GET("/users/:userId", controllers.GetUserFromAPI)
	router.GET("/results/:userId", controllers.GetResultFromAPI)
	router.GET("/results/:userId/wg", controllers.GetResultWGFromAPI)
	router.GET("/results/:userId/ch", controllers.GetResultChanFromAPI)
	router.GET("/sites/:siteId", controllers.GetSiteFromAPI)
	router.GET("/countries/:countryId", controllers.GetCountryFromAPI)

	router.Run(port)
}
