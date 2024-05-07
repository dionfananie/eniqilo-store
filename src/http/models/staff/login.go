package user

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=10,max=16"`
	Password    string `json:"password" binding:"required,min=5,max=15"`
}
