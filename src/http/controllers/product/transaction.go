package productController

import (
	"eniqilo-store/src/http/models/product"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
	} else {
		limitQuery = "LIMIT 5"
	}

	if offset := c.Query("offset"); offset != "" {
		offsetQuery = fmt.Sprintf("OFFSET $%d", len(params)+1)
		params = append(params, offset)
	}

	if createdAt := c.Query("createdAt"); createdAt != "" {
		if createdAt == "desc" {
			createdAt = "DESC"
		} else {
			createdAt = "ASC"
		}
		conditionOrders = append(conditionOrders, fmt.Sprintf("created_at %s", createdAt))
	} else {
		conditionOrders = append(conditionOrders, "created_at DESC")

	}

	if len(conditionOrders) > 0 {
		orderByQuery = "ORDER BY " + strings.Join(conditionOrders, ", ")
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	if orderByQuery != "" {
		baseQuery += " " + orderByQuery
	}

	if limitQuery != "" {
		baseQuery += " " + limitQuery
	}

	if offsetQuery != "" {
		baseQuery += " " + offsetQuery
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

	transactionIds := make([]string, 0, len(transactions))
	for _, trx := range transactions {
		transactionIds = append(transactionIds, trx.TransactionId)
	}

	itemRows, err := dbase.DB.Query("SELECT id, transaction_id, product_id, quantity FROM transaction_items WHERE transaction_id = ANY($1)", pq.Array(transactionIds))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer itemRows.Close()

	var transactionItems = make([]product.ProductTransactionItemModel, 0, 5)
	for itemRows.Next() {

		var item product.ProductTransactionItemModel
		if err := itemRows.Scan(&item.Id,
			&item.TransactionId,
			&item.ProductId,
			&item.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		transactionItems = append(transactionItems, item)
	}

	transactionItemsMap := make(map[string][]product.ProductTransactionItemModel)
	for _, item := range transactionItems {
		transactionItemsMap[item.TransactionId] = append(transactionItemsMap[item.TransactionId], item)
	}

	for i, transaction := range transactions {
		items, ok := transactionItemsMap[transaction.TransactionId]
		if !ok {
			continue
		}

		for _, item := range items {
			quantity, err := strconv.Atoi(item.Quantity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			detail := product.ProductCheckoutDetail{
				ProductId: item.ProductId,
				Quantity:  quantity,
			}
			transactions[i].ProductDetails = append(transactions[i].ProductDetails, detail)

		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": transactions})

}
