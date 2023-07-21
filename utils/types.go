package utils

import "github.com/google/uuid"

type Role string
type Action string
type HttpMethod string
type HttpRoute string
type Resource string

type Credential struct {
	Role     Role         `json:"role"`
	Resource Resource     `json:"resource"`
	Actions  []HttpMethod `json:"actions"`
}

type DtoCurrentUser struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid"`
	Email    string    `json:"email" validate:"required,email,lte=255"`
	RoleID   uuid.UUID `json:"role_id" validate:"required,uuid"`
	RoleName string    `json:"role_name" validate:"required"`
}
