package request

type Message struct {
	Prompt string `json:"prompt" binding:"required"`
}

func (message Message) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"propmpt.required": "propmpt is required",
	}
}

type QueryMessage struct {
	CID string `json:"c_id" binding:"required"`
}

func (queryMessage QueryMessage) GetQueryMessage() ValidatorMessages {
	return ValidatorMessages{
		"c_id.required": "c_id is required",
	}
}
