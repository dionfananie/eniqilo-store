package customerController

import (
	"log"
	"net/http"
	"regexp"

	customerModel "eniqilo-store/src/http/models/customer"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (dbase *V1Customer) CustomerRegister(c *gin.Context) {
	var req customerModel.CustomerRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var phoneExist bool
	err := dbase.DB.QueryRow("SELECT EXISTS(SELECT 1 from customers WHERE phone_number = $1)", req.PhoneNumber).Scan(&phoneExist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if phoneExist {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone Number already exists"})
		return
	}

	re := regexp.MustCompile(`^\+(?:[0-9]-? ?){6,14}[0-9]$`)
	if !re.MatchString(req.PhoneNumber) {
		log.Println("Phone number is not valid:", req.PhoneNumber)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone Number must contains country code"})
		return
	}

	var UserId string
	err = dbase.DB.QueryRow("INSERT INTO customers (name, phone_number) VALUES ($1, $2) RETURNING id", req.Name, req.PhoneNumber).Scan(&UserId)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "data": gin.H{
		"userId":      UserId,
		"phoneNumber": req.PhoneNumber,
		"name":        req.Name,
	}})

}
