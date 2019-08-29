package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"../utils"
)

var (
	paramUserID="userId"
)

func GetUserFromAPI(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param(paramUserID))
	if err != nil {
		apierr := utils.Apierror{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apierr.Status, apierr)
		return
	}

	response, apierr := services.GetUser(userId)
	if err != nil {
		c.JSON(apierr.Status, apierr)
		return
	}

	c.JSON(http.StatusOK, response)
}
