package v1routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type V1Routes struct {
	Gin *gin.RouterGroup
	DB  *sql.DB
}

type iV1Routes interface {
	MountStaff()
	MountCustomer()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
