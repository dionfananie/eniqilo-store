package customerController

import (
	customerModel "eniqilo-store/src/http/models/customer"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (dbase *V1Customer) CustomerList(c *gin.Context) {
	baseQuery := "SELECT id, name, phone_number from customers WHERE TRUE"
	var params []interface{}
	var conditions []string

	if phone := c.Query("phoneNumber"); phone != "" {
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(params)+1))
		params = append(params, phone)
	}

	if name := c.Query("name"); name != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE $%d", len(params)+1))
		params = append(params, "%"+name+"%")
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

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

	customers := make([]customerModel.CustomerList, 0)

	for rows.Next() {
		var customerItem customerModel.CustomerList
		if err := rows.Scan(&customerItem.Id, &customerItem.Name, &customerItem.PhoneNumber); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customerItem)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": customers})
}
