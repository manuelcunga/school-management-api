package dtos

type CreateSchoolInput struct {
	Name       string `json:"name" `
	Email      string `json:"email" `
	RoomNumber string `json:"room_number" `
	Province   string `json:"province" `
}
