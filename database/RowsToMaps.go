package database

import (
	"database/sql"
	"go.uber.org/zap"
	"strings"
)

func RowsToMaps(rows *sql.Rows) []*map[string]string {
	result := make([]*map[string]string, 0, 5)
	cols, _ := rows.Columns()
	colLen := len(cols)
	colNames := make([]string, colLen, colLen)
	toCamelCase(colNames, cols)
	for rows.Next() {
		columns := make([]string, colLen)
		columnPointers := make([]interface{}, colLen)
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			zap.L().Error(err.Error())
		}
		m := make(map[string]string)
		for i, _ := range cols {
			val := columnPointers[i].(*string)
			m[colNames[i]] = *val
		}
		result = append(result, &m)
	}
	return result
}

func toCamelCase(colNames []string, cols []string) {
	for i, colName := range cols {
		underLineIndex := strings.Index(colName, "_")
		if underLineIndex > 0 {
			colName = colName[0:underLineIndex] + strings.ToUpper(string(colName[underLineIndex+1])) + colName[underLineIndex+2:]
		}
		colNames[i] = colName
	}
}
