package apperror

import "encoding/json"

type AppError struct {
	Err        error  `json:"-"`
	Message    string `json:"message,omitempty"`
	DevMessage string `json:"dev_message,omitempty"`
	Code       string `json:"code,omitempty"`
}

func NewAppError(err error, m, dm, c string) *AppError {
	return &AppError{
		Err:        err,
		Message:    m,
		DevMessage: dm,
		Code:       c,
	}
}

func NewSystemError(err error) *AppError {
	return NewAppError(err, "internal service error", err.Error(), "ES-5-0") // code:  ES - event service
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return data
}
