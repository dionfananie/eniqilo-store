package product

type ProductRegisterModel struct {
	Name        string `json:"name" binding:"required,min=1,max=50"`
	Sku         string `json:"sku" binding:"required,min=1,max=30"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"imageUrl" binding:"required"`
	Notes       string `json:"notes" binding:"required,min=1,max=200"`
	Price       int64  `json:"price" binding:"required"`
	Stock       int64  `json:"stock" binding:"required,min=1"`
	Location    string `json:"location" binding:"required,min=1,max=200"`
	IsAvailable bool   `json:"isAvailable"`
}
