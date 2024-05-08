package staffController

import (
	"net/http"

	"eniqilo-store/src/helpers/jwt"
	"eniqilo-store/src/helpers/password"
	"eniqilo-store/src/http/models/staff"

	"github.com/gin-gonic/gin"
)

func (dbase *V1Staff) StaffLogin(c *gin.Context) {
	var req staff.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Name string
	var phoneNumber string
	var Pass string
	var UserId string

	err := dbase.DB.QueryRow("SELECT id, name, password, phone_number from staffs WHERE phone_number = $1", req.PhoneNumber).Scan(&UserId, &Name, &Pass, &phoneNumber)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return

	}

	errs := password.Verify(Pass, req.Password)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request doesnt pass validation"})
		return

	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User logged successfully", "data": gin.H{
		"userId":      UserId,
		"phoneNumber": phoneNumber,
		"name":        Name,
		"accessToken": accessToken,
	}})

}
