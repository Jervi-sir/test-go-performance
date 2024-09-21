package handlers

import (
	"fmt"
	"log"
	"net/http"
	"octa_api_go/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchItems(c *gin.Context) {
	// Default values for pagination
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	// Calculate offset for pagination
	limit := 10
	offset := (page - 1) * limit

	// SQL query to fetch 10 random items
	query := fmt.Sprintf(`SELECT id, name, details, price, sizes, genders, images, last_reposted, wilaya_code
		FROM items
		ORDER BY RANDOM()
		LIMIT %d OFFSET %d`, limit, offset)

	// Log the SQL query for debugging
	// log.Println("Executing SQL Query:", query)

	// Execute the query
	rows, err := database.DBPool.Query(c.Request.Context(), query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Get the column descriptions (only the names)
	columnDescriptions := rows.FieldDescriptions()
	columnNames := make([]string, len(columnDescriptions))
	for i, col := range columnDescriptions {
		columnNames[i] = string(col.Name)
	}

	// Dynamically handle the results using []interface{}
	var items []map[string]interface{}

	for rows.Next() {
		// Create a slice to store the row values
		values := make([]interface{}, len(columnNames))
		valuePtrs := make([]interface{}, len(columnNames))

		// Create pointers to store the values
		for i := range columnNames {
			valuePtrs[i] = &values[i]
		}

		// Scan the row into the value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row", "details": err.Error()})
			return
		}

		// Map the values to column names
		rowMap := make(map[string]interface{})
		for i, colName := range columnNames {
			rowMap[colName] = values[i]
		}

		// Append the row map to the items slice
		items = append(items, rowMap)
	}

	// Return the results as JSON
	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"items": items,
	})
}
