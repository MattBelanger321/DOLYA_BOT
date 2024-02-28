package utilities

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Creates a `*bytes.Buffer` for a message, to be handled by an `*http.Request`.
//
// Can also be used in the `CreateRequest` function.
func CreateMessage(preview_url bool, body string) *bytes.Buffer {
	message := WhatsAppMessage{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               To,
		Type:             "text",
		Text: TextStruct{
			PreviewURL: true,
			Body:       "DOLYA BOT ACTIVATED",
		},
	}

	// Turn the message struct into a JSON object.
	jsonData, err := json.Marshal(message)
	CheckError("Error marshalling the JSON", err)

	return bytes.NewBuffer(jsonData)
}

// Creates an `*http.Request` for sending a message.
func CreateRequest(method string, data *bytes.Buffer) *http.Request {
	req, err := http.NewRequest(method, URL, data)
	CheckError("Error creating request", err)

	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Content-Type", "application/json")

	return req
}

// Combines the functionality of `CreateMessage`, `CreateRequest` and `SendMessage`
// to create and send the message.
func CreateAndSendMessage(preview_url bool, body, method string) {
	data := CreateMessage(preview_url, body)
	req := CreateRequest(method, data)
	SendMessage(req)
}
