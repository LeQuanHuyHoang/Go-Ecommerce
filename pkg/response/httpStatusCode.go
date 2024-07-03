package response

const (
	ErrCodeSuccess = 20001 // Success
	ErrCodeInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // Token is invalid
)

var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeInvalid: "invalid",

	ErrInvalidToken: "token is invalid",
}
