package errs

import "net/http"

type AppErr struct {
	Code int
	Msg  string
}

func (ae AppErr) AsMsg() *AppErr {
	return &AppErr{Msg: ae.Msg}
}

func NewNotFoundErr(msg string) *AppErr {
	return &AppErr{Msg: msg, Code: http.StatusNotFound}
}

func NewValidationErr(msg string) *AppErr {
	return &AppErr{Msg: msg, Code: http.StatusUnprocessableEntity}
}

func NewUnExpectedErr(msg string) *AppErr {
	return &AppErr{Msg: msg, Code: http.StatusInternalServerError}
}

func NewAuthenticationErr(msg string) *AppErr {
	return &AppErr{Msg: msg, Code: http.StatusUnauthorized}
}

func NewAuthorizationErr(msg string) *AppErr {
	return &AppErr{Msg: msg, Code: http.StatusForbidden}
}
