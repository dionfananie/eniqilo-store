package v1routes

import (
	productController "eniqilo-store/src/http/controllers/product"
	middleware "eniqilo-store/src/http/middlewares"
)

func (i *V1Routes) MountProduct() {
	g := i.Gin.Group("/product")

	productController := productController.New(&productController.V1User{
		DB: i.DB,
	})
	g.Use(middleware.AuthMiddleware())
	g.POST("/", productController.ProductRegister)
	g.GET("/", productController.ProductList)
	g.GET("/customer", productController.ProductListCustomer)
	g.GET("/checkout", productController.ProductCheckout)

}
