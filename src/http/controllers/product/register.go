package productController

import (
	"net/http"
	"regexp"

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

	if category := req.Category; category != "" {
		err := validation.ValidateCategory(category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

	re := regexp.MustCompile(`[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	if !re.MatchString(req.ImageUrl) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image Url is not valid"})
		return
	}

	err := dbase.DB.QueryRow("INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at", req.Name, req.Sku, req.Category, req.ImageUrl, req.Notes, req.Price, req.Stock, req.Location, req.IsAvailable).Scan(&ProductId, &CreatedAt)

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
