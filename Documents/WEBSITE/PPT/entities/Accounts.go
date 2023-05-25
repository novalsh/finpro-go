package entities

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Accounts struct {
	ID       int64           `json:"id"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Phone    string          `json:"phone"`
	Role     string          `json:"role"`
	Token    string          `json:"token"`
	CreatedAt *timestamp.Timestamp `json:"createdAt"`
	UpdatedAt *timestamp.Timestamp `json:"updatedAt"`
}
