package v1routes

import (
	staffController "eniqilo-store/src/http/controllers/staff"
)

func (i *V1Routes) MountStaff() {
	g := i.Gin.Group("/staff")

	staffController := staffController.New(&staffController.V1Staff{
		DB: i.DB,
	})

	g.POST("/register", staffController.StaffRegister)
	g.POST("/login", staffController.StaffLogin)

}
