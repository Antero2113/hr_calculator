package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetTable(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT 
				e.name,
				CONCAT(e.region, ' ', e.department),
				e.position,
				jr.client,
				pp.name,
				pp.operations,
				pp.measure,
				pp.min,
				pp.max,
				pp.period_type,
				pp.period_count
			FROM job_records jr
			LEFT JOIN employees e ON jr.employee = e.id
			LEFT JOIN processes p ON jr.process = p.id
			LEFT JOIN process_processes_properties ppp ON p.id = ppp.process_id
			LEFT JOIN process_properties pp ON ppp.property_id = pp.id
		`)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		defer rows.Close()

		var result []map[string]interface{}

		for rows.Next() {
			var (
				name, department, position, client,
				processName, operations, measure, periodType string
				min, max, periodCount                         int
			)

			rows.Scan(
				&name, &department, &position, &client,
				&processName, &operations, &measure,
				&min, &max, &periodType, &periodCount,
			)

			row := map[string]interface{}{
				"employee_name": name,
				"department":    department,
				"position":      position,
				"client":        client,
				"process_name":  processName,
				"operations":    operations,
				"measure":       measure,
				"min":           min,
				"max":           max,
				"period_type":   periodType,
				"period_count":  periodCount,
			}

			result = append(result, row)
		}

		json.NewEncoder(w).Encode(result)
	}
}