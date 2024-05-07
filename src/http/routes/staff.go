package v1routes

import (
	staffController "eniqilo-store/src/http/controllers/staff"
)

func (i *V1Routes) MountStaff() {
	g := i.Gin.Group("/staff")

	staffController := staffController.New(&staffController.V1User{
		DB: i.DB,
	})

	g.POST("/register", staffController.Register)
	g.GET("/login", staffController.Login)

}
