package staff

type RegisterRequest struct {
	Name        string `json:"name" binding:"required,min=5,max=50"`
	PhoneNumber string `json:"phoneNumber" binding:"required,min=10,max=16"`
	Password    string `json:"password" binding:"required,min=5,max=15"`
}
