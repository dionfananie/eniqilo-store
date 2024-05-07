package v1controller

import (
	"net/http"

	"eniqilo-store/src/helpers"
	"eniqilo-store/src/http/models/user"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (dbase *V1User) Login(c gin.Context) (err error) {
	var req user.LoginRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Name string
	var phone_number string
	var Pass string
	var UserId uint64

	err = dbase.DB.QueryRow("SELECT id, name, password, phone_number from users WHERE phone_number = $1", req.phone_number).Scan(&UserId, &Name, &Pass, &phone_number)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return

	}

	errs := helpers.VerifyPassword(Pass, req.Password)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request doesnt pass validation"})
		return

	}

	accessToken := helpers.Generate(&helpers.TokenPayload{
		UserId: UserId,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User logged successfully", "data": gin.H{
		"phone_number": phone_number,
		"name":         Name,
		"accessToken":  accessToken,
	}})
}

func (dbase *V1User) Register(c gin.Context) (err error) {
	var req user.RegisterRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := helpers.Hash(req.Password)

	var phoneExist bool
	err = dbase.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = $1)", req.PhoneNumber).Scan(&phoneExist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if phoneExist {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone Number already exists"})
		return
	}

	var UserId uint64
	err = dbase.DB.QueryRow("INSERT INTO users (name, phone_number, password) VALUES ($1, $2, $3) RETURNING id", req.Name, req.phone_number, hashedPassword).Scan(&UserId)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken := helpers.Generate(&helpers.TokenPayload{
		UserId: UserId,
	})

	// return c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "data": gin.H{
	// 	"phoneNumber": req.PhoneNumber,
	// 	"name":        req.Name,
	// 	"accessToken": accessToken,
	// }})

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}
