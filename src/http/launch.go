package http

import (
	v1routes "eniqilo-store/src/http/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (i *Http) Launch() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	basePath := "/v1"
	v1 := v1routes.New(&v1routes.V1Routes{
		Gin: r.Group(basePath),
		DB:  i.DB,
	})
	v1.MountUser()
	// router := r.Group("/v1")
	// router.POST("/user/login", controller.Login)
	// router.POST("/user/register", controller.Register)

	r.Run(":8080")
}
