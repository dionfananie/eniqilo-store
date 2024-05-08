package staffController

import (
	"net/http"

	"eniqilo-store/src/helpers/jwt"
	"eniqilo-store/src/helpers/password"
	"eniqilo-store/src/http/models/staff"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (dbase *V1Staff) StaffRegister(c *gin.Context) {
	var req staff.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := password.Hash(req.Password)

	var phoneExist bool
	err := dbase.DB.QueryRow("SELECT EXISTS(SELECT 1 from staffs WHERE phone_number = $1)", req.PhoneNumber).Scan(&phoneExist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if phoneExist {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone Number already exists"})
		return
	}

	var UserId string
	err = dbase.DB.QueryRow("INSERT INTO staffs (name, phone_number, password) VALUES ($1, $2, $3) RETURNING id", req.Name, req.PhoneNumber, hashedPassword).Scan(&UserId)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
	})

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "data": gin.H{
		"userId":      UserId,
		"phoneNumber": req.PhoneNumber,
		"name":        req.Name,
		"accessToken": accessToken,
	}})

}
