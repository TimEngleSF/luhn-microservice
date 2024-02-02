package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Digits struct {
	D int `json:"digits"`
}

func handleRequest(c *gin.Context) {
	var digitData Digits

	if err := c.ShouldBindJSON(&digitData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "must include digits"})
		return
	}

	if err := validate.Struct(digitData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorr": err.Error()})
		return
	}


	isValid := validateCard(digitData.D)

	c.JSON(http.StatusAccepted, gin.H{"valid" : isValid})	
	
}
