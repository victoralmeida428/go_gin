package dto

import (
	"abramed_go/dto/req"
	"abramed_go/dto/response"
)

type Responses struct {
	Login response.LoginResponse `json:"login"`
}

type Requests struct {
	Login req.LoginRequest `json:"login"`
}
