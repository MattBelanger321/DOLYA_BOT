package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mattbelanger321/dolya-bot/utilities"
)

var (
	url = "https://graph.facebook.com/v17.0/233298403205906/messages"
	// Replace with your actual token
	token = "EAALMF3jOoVgBOzwXHcNrEEbTRoewuJjc8hu24dH0xkmliiVz8kYSjpiY4j5YvGBdpcTuwtNSgokBGei8ZBM1c4B2tjD3z2JlGdJLbZA1nhpxs87PlrmsUL2e0pUiGXYAOsZA2KP4a2AEOk9oejM9GNfqFiARbTcgF3ao8rGxdhQmGIsHp0ZAeKQU9QuXugGIPCVQPw1szA1iXmbAkPYZD"
	// Replace with the recipient's phone number
	to = "+15199912176"
)

func main() {

	message := utilities.WhatsAppMessage{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "text",
		Text: utilities.TextStruct{
			PreviewURL: true,
			Body:       "DOLYA BOT ACTIVATED",
		},
	}

	// Turn the message struct into a JSON object.
	jsonData, err := json.Marshal(message)
	utilities.CheckError("Error marshalling the JSON", err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	utilities.CheckError("Error creating request", err)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	utilities.CheckError("Error making request", err)
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	// You can handle the response body as needed
	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(responseBody))
}
