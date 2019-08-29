package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResultFromAPI(c *gin.Context) {

	if utils.CircuitBreaker.State != "CLOSE" {
		c.JSON(500, utils.Apierror{
			Message: "Server Unavailable",
			Status:  500,
		})
		return
	}

	userId, err := strconv.Atoi(c.Param(paramUserID))
	if err != nil {
		apierr := utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apierr.Status, apierr)
		return
	}
	for i := 0; i < utils.CircuitBreaker.CantErrors; i++ {
		response, apierr := services.GetResult(userId)
		if apierr != nil {
			if apierr.Status == 533 {
				continue
			}
			c.JSON(apierr.Status, apierr)
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(500, utils.Apierror{
		Message: "Server Unavailable",
		Status:  500,
	})

	go utils.TimeOut()

}
