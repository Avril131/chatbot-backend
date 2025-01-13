package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{ErrorCode: 40000, ErrorMsg: "BusinessError"},
	ValidateError: CustomError{ErrorCode: 42200, ErrorMsg: "ValidateError"},
	TokenError:    CustomError{ErrorCode: 40100, ErrorMsg: "Failed login authorization"},
}
