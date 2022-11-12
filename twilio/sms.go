package twilio

import (
	"encoding/json"
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
	params.SetTo(toNumber)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
		return true
	}
}
