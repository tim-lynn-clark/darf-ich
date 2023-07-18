package utils

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
