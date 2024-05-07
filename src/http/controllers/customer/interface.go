package customerController

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type V1User struct {
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
	CustomerList(c *gin.Context)
	CustomerRegister(c *gin.Context)
}

func New(v1User *V1User) iV1User {
	return v1User
}
