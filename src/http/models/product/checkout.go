package product

type ProductCheckoutModel struct {
	CustomerId     string   `json:"customerId" binding:"required"`
	ProductDetails []detail `json:"productDetails" binding:"required"`
	Paid           int      `json:"paid" binding:"required"`
	Change         int      `json:"change" binding:"required"`
}

type detail struct {
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}
