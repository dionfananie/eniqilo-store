package productController

import (
	"eniqilo-store/src/http/models/product"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (dbase *V1Product) ProductTransactions(c *gin.Context) {
	baseQuery := "SELECT id, customer_id, paid, change, created_at FROM transactions WHERE TRUE"
	var params []interface{}
	var conditions, conditionOrders []string
	var limitQuery, offsetQuery, orderByQuery string

	if customerId := c.Query("customerId"); customerId != "" {
		conditions = append(conditions, fmt.Sprintf("customer_id = $%d", len(params)+1))
		params = append(params, customerId)
	}

	if limit := c.Query("limit"); limit != "" {
		limitQuery = fmt.Sprintf("LIMIT $%d", len(params)+1)
		params = append(params, limit)
	}

	if offset := c.Query("offset"); offset != "" {
		offsetQuery = fmt.Sprintf("OFFSET $%d", len(params)+1)
		params = append(params, offset)
	}

	if createdAt := c.Query("created_at"); createdAt != "" {
		if createdAt == "desc" {
			createdAt = "DESC"
		} else {
			createdAt = "ASC"
		}
		conditionOrders = append(conditionOrders, fmt.Sprintf("created_at %s", createdAt))
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

	rows, err := dbase.DB.Query(baseQuery, params...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	transactions := make([]product.ProductTransactionModel, 0, 20)
	for rows.Next() {
		var transaction product.ProductTransactionModel
		if err := rows.Scan(&transaction.TransactionId,
			&transaction.CustomerId,
			&transaction.Paid,
			&transaction.Change,
			&transaction.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": transactions})

}
