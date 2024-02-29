# DOLYA_BOT
This is the backend for the DOLYA bot.

In order to get the bot working after forking this repository, create the file `utilities/ids.go` with the following content:
```golang
package utilities

var (
	URL     = "https://graph.facebook.com/v17.0/"
	ID      = "SOME_NUMERIC_ID/"
	Command = "messages/"
	Token   = "SOME_TOKEN"
	To      = "SOME_NUMBER" // You may need to include a `1` in front of your phone number.
)
```