package product

type ProductTransactionModel struct {
	TransactionId  string                  `json:"transactionId"`
	CustomerId     string                  `json:"customerId"`
	ProductDetails []ProductCheckoutDetail `json:"productDetails"`
	Paid           int                     `json:"paid"`
	Change         int                     `json:"change"`
	CreatedAt      string                  `json:"createdAt"`
}
