package auth_const

import "fmt"

type Preload struct {
	Role   string `json:"role"`
	UserId int64  `json:"user_id"`
}

func (p Preload) GetAuthId() string {

	return fmt.Sprintf("auth.%d", p.UserId)
}
