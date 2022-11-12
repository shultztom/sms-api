package twilio

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

func SendSMS(toNumber string, body string) bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(os.Getenv("FROM_NUMBER"))
	// Optional: Used for monitoring within twilio
	if os.Getenv("MESSAGING_SERVICE_SID") != "" {
		params.SetMessagingServiceSid(os.Getenv("MESSAGING_SERVICE_SID"))
	}
	params.SetTo(toNumber)
	params.SetBody(body)

	_, err = client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}
