package utilities

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Creates a `*bytes.Buffer` for a message, to be handled by an `*http.Request`.
//
// Can also be used in the `CreateRequest` function.
func CreateMessage(preview_url bool, phone, body string) *bytes.Buffer {
	message := WhatsAppMessage{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone,
		Type:             "text",
		Text: TextStruct{
			PreviewURL: false,
			Body:       body,
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
func CreateAndSendMessage(preview_url bool, phone, body, method string) {
	data := CreateMessage(preview_url, phone, body)
	req := CreateRequest(method, data)
	SendMessage(req)
}
