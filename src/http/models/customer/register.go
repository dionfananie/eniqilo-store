package customerModel

type CustomerRegister struct {
	Name        string `json:"name" binding:"required,min=5,max=50"`
	PhoneNumber string `json:"phoneNumber" binding:"required,min=10,max=16"`
}
