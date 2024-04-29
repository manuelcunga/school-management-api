package dtos

import "time"

type SchoolOutput struct {
	ID         string    `json:"id" `
	Name       string    `json:"name" `
	Email      string    `json:"email" `
	RoomNumber string    `json:"Room_number" `
	Province   string    `json:"province" `
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" `
}
