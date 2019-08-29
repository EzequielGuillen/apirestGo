package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
)

const (
	paramSiteID = "siteId"
)

func GetSiteFromAPI(c *gin.Context)  {

	siteid:=c.Param(paramSiteID)

	response, err := services.GetSite(siteid)
	if err != nil {
		c.JSON(err.Status,err)
		return
	}

	c.JSON(http.StatusOK,response)
}
