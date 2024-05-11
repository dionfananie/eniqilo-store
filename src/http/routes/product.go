package v1routes

import (
	productController "eniqilo-store/src/http/controllers/product"
	middleware "eniqilo-store/src/http/middlewares"
)

func (i *V1Routes) MountProduct() {
	g := i.Gin.Group("/product")

	productController := productController.New(&productController.V1Product{
		DB: i.DB,
	})
	g.GET("/customer", productController.ProductList)

	g.Use(middleware.AuthMiddleware())
	g.POST("/", productController.ProductRegister)
	g.GET("/", productController.ProductList)
	g.PUT("/:id", productController.ProductEdit)
	g.DELETE("/:id", productController.ProductDelete)
	g.GET("/checkout/history", productController.ProductTransactions)
	g.POST("/checkout", productController.ProductCheckout)

}
