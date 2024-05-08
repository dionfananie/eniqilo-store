package productController

import (
	"net/http"

	"eniqilo-store/src/helpers/validation"
	"eniqilo-store/src/http/models/product"

	"github.com/gin-gonic/gin"

	"github.com/lib/pq"
)

func (dbase *V1Product) ProductRegister(c *gin.Context) {
	var req product.ProductRegisterModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ProductId string
	var CreatedAt string

	isAvailable := false
	if req.Stock > 1 {
		isAvailable = true
	}

	if category := req.Category; category != "" {
		err := validation.ValidateCategory(category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
	err := dbase.DB.QueryRow("INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at", req.Name, req.Sku, req.Category, req.ImageUrl, req.Notes, req.Price, req.Stock, req.Location, isAvailable).Scan(&ProductId, &CreatedAt)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "successfully add product", "data": gin.H{
		"id":        ProductId,
		"createdAt": CreatedAt,
	}})

}
