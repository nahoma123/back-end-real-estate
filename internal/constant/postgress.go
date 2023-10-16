package constant

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

/*
Filters
 1. by gte for date, by gte for age [gte]
 2. equals [=]
 3. contains [contains]
 4. not equals [!=]
*/

func GetResults(ctx context.Context, db *gorm.DB, tableName string, filterPagination *FilterPagination, results *[]map[string]interface{}) error {
	// Create a GORM DB instance for the specified table
	table := db.Table(tableName)

	// Apply filters
	for _, f := range filterPagination.Filters {
		if f.Operator == "gte" {
			table = table.Where(fmt.Sprintf("%s >= ?", f.Field), f.Value)
		} else if f.Operator == "=" {
			table = table.Where(fmt.Sprintf("%s = ?", f.Field), f.Value)
		} else if f.Operator == "contains" {
			table = table.Where(fmt.Sprintf("%s ILIKE ?", f.Field), fmt.Sprintf("%%%s%%", f.Value))
		} else if f.Operator == "!=" {
			table = table.Where(fmt.Sprintf("%s != ?", f.Field), f.Value)
		} else {
			// Handle other operators
		}
	}

	// Apply pagination
	table = table.Offset((filterPagination.Pagination.Page - 1) * filterPagination.Pagination.PerPage).Limit(filterPagination.Pagination.PerPage)

	// Execute the query and retrieve the results
	if err := table.Find(results).Error; err != nil {
		return err
	}

	// Get the total count of documents that match the filter
	var totalCount int64
	if err := table.Count(&totalCount).Error; err != nil {
		return err
	}

	filterPagination.TotalCount = totalCount
	filterPagination.TotalPages = int(totalCount/int64(filterPagination.Pagination.PerPage) + 1)

	return nil
}
