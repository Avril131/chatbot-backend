package request

type Login struct {
	Token string `json:"token" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"token.required": "token is required",
	}
}
