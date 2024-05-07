package v1routes

import (
	customerController "eniqilo-store/src/http/controllers/customer"
)

func (i *V1Routes) MountCustomer() {
	g := i.Gin.Group("/customer")

	customerController := customerController.New(&customerController.V1User{
		DB: i.DB,
	})

	g.POST("/register", customerController.Register)

}
