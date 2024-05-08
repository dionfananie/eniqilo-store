package v1routes

import (
	customerController "eniqilo-store/src/http/controllers/customer"
	middleware "eniqilo-store/src/http/middlewares"
)

func (i *V1Routes) MountCustomer() {
	g := i.Gin.Group("/customer")

	customerController := customerController.New(&customerController.V1Customer{
		DB: i.DB,
	})
	g.Use(middleware.AuthMiddleware())
	g.POST("/register", customerController.CustomerRegister)
	g.GET("/", customerController.CustomerList)

}
