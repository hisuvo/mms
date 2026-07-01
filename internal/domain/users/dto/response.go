package dto

import "time"

type ReqisterResponse struct {
	ID        uint `json:""`
	UseName   string `json:"user_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}