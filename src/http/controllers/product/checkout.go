package productController

import (
	"eniqilo-store/src/http/models/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (dbase *V1Product) ProductCheckout(c *gin.Context) {
	var req product.ProductCheckoutModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customerIdExist bool
	err := dbase.DB.QueryRow("SELECT EXIST(SELECT 1 from customers WHERE id = $1)", customerIdExist).Scan(&customerIdExist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !customerIdExist {
		c.JSON(http.StatusNotFound, gin.H{"message": "Customer id not found"})
	}

	productIds := make([]string, len(req.ProductDetails))
	for i, detail := range req.ProductDetails {
		productIds[i] = detail.ProductId
	}

	rows, err := dbase.DB.Query("SELECT id, name, is_available, category, sku, price, stock, image_url, location, created_at from products WHERE id = ANY($1)", productIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	countRows := 0
	products := make([]product.ProductListModel, len(productIds))
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
		c.JSON(http.StatusNotFound, gin.H{"message": "One of the product id is not found"})
	}

	productStockChanges := make([]product.ProductStockChangesModel, len(productIds))
	for _, productItem := range products {
		var quantity int
		for _, detail := range req.ProductDetails {
			if detail.ProductId == productItem.Id {
				quantity = detail.Quantity
				break
			}
		}
		productStockChanges = append(productStockChanges, product.ProductStockChangesModel{
			Id:         productItem.Id,
			FinalStock: int(productItem.Stock) - quantity,
		})

		if int(productItem.Stock)-quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "One of the product stock is not enough"})
			return
		}
	}

	for _, productChange := range productStockChanges {
		_, err := dbase.DB.Exec("UPDATE products SET stock = $1 WHERE id = $2", productChange.FinalStock, productChange.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//TODO: SAVE THE CHECKOUT DATA

	c.JSON(http.StatusCreated, gin.H{"message": "successfully add product", "data": gin.H{}})

}
