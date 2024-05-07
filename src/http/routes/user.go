package v1routes

import (
	v1controller "eniqilo-store/src/http/controllers"
)

func (i *V1Routes) MountUser() {
	g := i.Gin.Group("/user")

	userController := v1controller.New(&v1controller.V1User{
		DB: i.DB,
	})

	g.POST("/register", userController.Register)
	g.POST("/login", userController.Login)

}
