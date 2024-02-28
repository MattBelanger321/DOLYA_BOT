package utilities

import (
	"fmt"
	"io"
	"net/http"
)

// Sends a message encoded inside an `*http.Request` object.
//
// Further handles the response status.
func SendMessage(req *http.Request) {
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError("Error making request", err)
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	responseBody, err := io.ReadAll(resp.Body)
	CheckError("Error reading response body", err)
	fmt.Println("Response Body:", string(responseBody))
}
