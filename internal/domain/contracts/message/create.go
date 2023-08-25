package contracts

type CreateMessageRequest struct {
	Text string `json:"text"`
}

type CreateMessageResponse struct {
	MessageId string `json:"messageId"`
}
