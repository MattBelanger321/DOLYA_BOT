package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://graph.facebook.com/v17.0/233298403205906/messages"
	token := "EAALMF3jOoVgBO94ufYYMxnT49tP05HaB9gf89DRvdMHzJy6X9QiWJFq0TWZBdgKtLUgH1NSarA9t6q75qP7xmV8VQO0oEbUKNZAcTNigA4PamWFQwt0nNCZCR5OKJAEa06Y6HYcOhLajAFABRInFUAsYb8swKJdXXAJi1CFgzrYs4h65rTM87ZAxCqNBkBGuChmc4OfgkG8Trze2WSeP" // Replace with your actual token
	to := "+15199912176"                                                                                                                                                                                                             // Replace with the recipient's phone number

	jsonData := []byte(`{
		"messaging_product": "whatsapp",
		"recipient_type": "individual",
		"to": "` + to + `",
		"type": "text",
		"text": {
			"preview_url": true,
			"body": "DOLYA BOT ACTIVATED"
		}
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	// You can handle the response body as needed
	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(responseBody))
}
