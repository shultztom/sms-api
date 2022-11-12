package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sms-api/twilio"
)

type smsRequest struct {
	ToNumber string `json:"toNumber"`
	Body     string `json:"body"`
}

func PostSMS(c *gin.Context) {
	var newRequest smsRequest

	if err := c.BindJSON(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse body!",
		})
		return
	}

	if newRequest.ToNumber == "" || newRequest.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Body!",
		})
		return
	}

	successfullySent := twilio.SendSMS(newRequest.ToNumber, newRequest.Body)
	if successfullySent {
		c.JSON(http.StatusOK, gin.H{
			"message": "SENT TEXT",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error sending SMS message",
		})
	}
}
