package product

type ProductCheckoutModel struct {
	CustomerId     string                  `json:"customerId" binding:"required"`
	ProductDetails []ProductCheckoutDetail `json:"productDetails" binding:"required"`
	Paid           int                     `json:"paid" binding:"required"`
	Change         int                     `json:"change"`
}

type ProductCheckoutDetail struct {
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}
