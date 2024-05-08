package product

type ProductListModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Category    string `json:"category"`
	ImageUrl    string `json:"imageUrl"`
	Notes       string `json:"notes"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Location    string `json:"location"`
	IsAvailable string `json:"isAvailable"`
	CreatedAt   string `json:"createdAt"`
}
