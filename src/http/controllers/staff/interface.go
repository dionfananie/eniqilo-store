package staffController

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type V1Staff struct {
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
	StaffRegister(c *gin.Context)
	StaffLogin(c *gin.Context)
}

func New(v1Staff *V1Staff) iV1User {
	return v1Staff
}
