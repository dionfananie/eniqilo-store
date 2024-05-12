package productController

import (
	"eniqilo-store/src/http/models/product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (dbase *V1Product) ProductCheckout(c *gin.Context) {
	var req product.ProductCheckoutModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := uuid.Parse(req.CustomerId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wrong customer id"})
		return
	}

	var customerIdExist bool
	err = dbase.DB.QueryRow("SELECT EXISTS(SELECT 1 from customers WHERE id = $1)", req.CustomerId).Scan(&customerIdExist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !customerIdExist {
		c.JSON(http.StatusNotFound, gin.H{"message": "Customer id not found"})
	}

	productIds := make([]string, len(req.ProductDetails))
	for i, detail := range req.ProductDetails {
		if string(detail.ProductId) == "null" || detail.ProductId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong product id"})
			return
		}

		_, err := uuid.Parse(detail.ProductId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wrong product id"})
			return
		}

		productIds[i] = detail.ProductId
	}

	rows, err := dbase.DB.Query("SELECT id, name, is_available, category, sku, price, stock, image_url, location, created_at from products WHERE id = ANY($1) AND is_available = true", pq.Array(productIds))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	countRows := 0
	products := make([]product.ProductListModel, 0, len(productIds))
	for rows.Next() {
		var productItem product.ProductListModel
		countRows++
		err := rows.Scan(
			&productItem.Id,
			&productItem.Name,
			&productItem.IsAvailable,
			&productItem.Category,
			&productItem.Sku,
			&productItem.Price,
			&productItem.Stock,
			&productItem.ImageUrl,
			&productItem.Location,
			&productItem.CreatedAt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		products = append(products, productItem)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if countRows != len(productIds) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "One of the product id is not found or not available"})
		return
	}

	productStockChanges := make([]product.ProductStockChangesModel, 0, len(productIds))
	var productTotalPrice int
	for _, productItem := range products {
		var quantity int
		for _, detail := range req.ProductDetails {
			if detail.ProductId == productItem.Id {
				quantity = detail.Quantity
				productTotalPrice = productTotalPrice + int(productItem.Price)*quantity
				break
			}
		}

		if int(productItem.Stock)-quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "One of the product stock is not enough"})
			return
		}
		productStockChanges = append(productStockChanges, product.ProductStockChangesModel{
			Id:         productItem.Id,
			FinalStock: int(productItem.Stock) - quantity,
		})

	}
	if req.Paid < productTotalPrice {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Paid not enough"})
		return
	}
	if req.Change != req.Paid-productTotalPrice || req.Change < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Change is wrong"})
		return
	}

	for _, productChange := range productStockChanges {
		_, err := dbase.DB.Exec("UPDATE products SET stock = $1 WHERE id = $2", productChange.FinalStock, productChange.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var transactionId string
	err = dbase.DB.QueryRow("INSERT INTO transactions (customer_id, paid, change) VALUES ($1, $2, $3) RETURNING id", req.CustomerId, req.Paid, req.Change).Scan(&transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	baseQuery := "INSERT INTO transaction_items (transaction_id, product_id, quantity) VALUES "
	for i, item := range req.ProductDetails {
		if i > 0 {
			baseQuery += ", "
		}
		baseQuery += fmt.Sprintf("('%s', '%s', %d)", transactionId, item.ProductId, item.Quantity)
	}
	err = dbase.DB.QueryRow(baseQuery).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product checkout success"})

}
