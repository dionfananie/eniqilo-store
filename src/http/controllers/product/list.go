package productController

import (
	"eniqilo-store/src/http/models/product"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (dbase *V1Product) ProductList(c *gin.Context) {
	baseQuery := "SELECT id, name, is_available, category, sku, price, stock, image_url, location, created_at from products WHERE TRUE"
	var params []interface{}
	var conditions, conditionOrders []string
	var limitQuery, offsetQuery, orderByQuery string

	if id := c.Query("id"); id != "" {
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(params)+1))
		params = append(params, id)
	}

	if name := c.Query("name"); name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(params)+1))
		params = append(params, name)
	}

	if isAvailable := c.Query("isAvailable"); isAvailable != "" {
		conditions = append(conditions, fmt.Sprintf("is_available = $%d", len(params)+1))
		params = append(params, isAvailable)
	}
	if category := c.Query("category"); category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(params)+1))
		params = append(params, category)
	}
	if sku := c.Query("sku"); sku != "" {
		println("len", len(params))
		conditions = append(conditions, fmt.Sprintf("sku = $%d", len(params)+1))
		params = append(params, sku)
	}

	if stock := c.Query("stock"); stock != "" {
		conditions = append(conditions, fmt.Sprintf("stock = $%d", len(params)+1))
		params = append(params, stock)
	}
	if location := c.Query("location"); location != "" {
		conditions = append(conditions, fmt.Sprintf("location = $%d", len(params)+1))
		params = append(params, location)
	}

	if limit := c.Query("limit"); limit != "" {
		limitQuery = fmt.Sprintf("LIMIT $%d", len(params)+1)
		params = append(params, limit)
	}
	if offset := c.Query("offset"); offset != "" {
		offsetQuery = fmt.Sprintf("OFFSET $%d", len(params)+1)
		params = append(params, offset)
	}

	if price := c.Query("price"); price != "" {
		if price == "desc" {
			price = "DESC"
		} else {
			price = "ASC"
		}
		conditionOrders = append(conditionOrders, fmt.Sprintf("price %s", price))

	} else {
		conditionOrders = append(conditionOrders, "price")
	}
	if createdAt := c.Query("created_at"); createdAt != "" {
		if createdAt == "desc" {
			createdAt = "DESC"
		} else {
			createdAt = "ASC"
		}
		conditionOrders = append(conditionOrders, fmt.Sprintf("created_at %s", createdAt))

	} else {
		conditionOrders = append(conditionOrders, "created_at")
	}

	if len(conditionOrders) > 0 {
		orderByQuery = "ORDER BY " + strings.Join(conditionOrders, ", ")
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	if limitQuery != "" {
		baseQuery += " " + limitQuery
	}

	if offsetQuery != "" {
		baseQuery += " " + offsetQuery
	}

	if orderByQuery != "" {
		baseQuery += " " + orderByQuery
	}

	println(baseQuery)
	println(params)
	rows, err := dbase.DB.Query(baseQuery, params...)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	products := make([]product.ProductListModel, 0)

	for rows.Next() {
		var productItem product.ProductListModel
		if err := rows.Scan(
			&productItem.Id,
			&productItem.Name,
			&productItem.IsAvailable,
			&productItem.Category,
			&productItem.Sku,
			&productItem.Price,
			&productItem.Stock,
			&productItem.ImageUrl,
			&productItem.Location,
			&productItem.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, productItem)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": products})
}
