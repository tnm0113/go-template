package user

import (
	"net/http"

	"github.com/c4i/go-template/internal/types"
)

var _ types.Response = (*UserResponse)(nil)

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func (res UserResponse) Code() int {
	return http.StatusOK
}

func (res UserResponse) Headers() map[string]string {
	return map[string]string{}
}

func (res UserResponse) Empty() bool {
	return false
}
