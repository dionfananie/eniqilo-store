package productController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (dbase *V1Product) ProductDelete(c *gin.Context) {
	id := c.Param("id")

	rows, err := dbase.DB.Exec("DELETE FROM products WHERE id = $1", id)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": " successfully delete product"})
}
