package utilities

type TextStruct struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type WhatsAppMessage struct {
	MessagingProduct string     `json:"messaging_product"`
	RecipientType    string     `json:"recipient_type"`
	To               string     `json:"to"`
	Type             string     `json:"type"`
	Text             TextStruct `json:"text"`
}
