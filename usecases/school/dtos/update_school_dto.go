package dtos

type UpdateSchoolInput struct {
	Name       string `json:"name" `
	Email      string `json:"email" `
	RoomNumber string `json:"Room_number" `
	Province   string `json:"province" `
}
