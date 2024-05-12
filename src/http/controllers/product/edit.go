package productController

import (
	"eniqilo-store/src/helpers/validation"
	"eniqilo-store/src/http/models/product"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (dbase *V1Product) ProductEdit(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error Id Product"})
		return
	}

	var req product.ProductRegisterModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter required"})
	}

	baseQuery := "UPDATE products SET "

	var params []interface{}
	var conditions []string
	if req.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(params)+1))
		params = append(params, req.Name)

	}
	if req.Sku != "" {
		conditions = append(conditions, fmt.Sprintf("sku = $%d", len(params)+1))
		params = append(params, req.Sku)

	}
	if req.Category != "" {

		err := validation.ValidateCategory(req.Category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(params)+1))
		params = append(params, req.Category)

	}
	if req.ImageUrl != "" {
		re := regexp.MustCompile(`[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

		if !re.MatchString(req.ImageUrl) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image Url is not valid"})
			return
		}
		conditions = append(conditions, fmt.Sprintf("image_url = $%d", len(params)+1))
		params = append(params, req.ImageUrl)

	}

	if req.Notes != "" {
		conditions = append(conditions, fmt.Sprintf("notes = $%d", len(params)+1))
		params = append(params, req.Notes)

	}
	if req.Price > 0 {
		conditions = append(conditions, fmt.Sprintf("price = $%d", len(params)+1))
		params = append(params, req.Price)
	}

	if req.Stock > 0 {
		conditions = append(conditions, fmt.Sprintf("stock = $%d", len(params)+1))
		params = append(params, req.Stock)
	}
	if req.Location != "" {
		conditions = append(conditions, fmt.Sprintf("location = $%d", len(params)+1))
		params = append(params, req.Location)
	}

	conditions = append(conditions, fmt.Sprintf("is_available = $%d", len(params)+1))
	params = append(params, req.IsAvailable)

	if len(conditions) > 0 {
		baseQuery = baseQuery + strings.Join(conditions, ", ")
	}
	whereQuery := fmt.Sprintf("WHERE id = '%s'", id)
	baseQuery = baseQuery + " " + whereQuery

	res, err := dbase.DB.Exec(baseQuery, params...)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully update cat"})
}
