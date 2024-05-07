package v1routes

import (
	"github.com/gin-gonic/gin"
)

type V1Routes struct {
	Gin *gin.Group
}

type iV1Routes interface {
	MountUser()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
