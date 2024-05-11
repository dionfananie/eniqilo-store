package productController

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type V1Product struct {
	DB *sql.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1User interface {
	ProductRegister(c *gin.Context)
	ProductList(c *gin.Context)
	ProductCheckout(c *gin.Context)
	ProductTransactions(c *gin.Context)
	ProductEdit(c *gin.Context)
	ProductDelete(c *gin.Context)
}

func New(v1Product *V1Product) iV1User {
	return v1Product
}
