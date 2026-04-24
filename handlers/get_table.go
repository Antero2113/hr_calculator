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
				e.position,
				jr.client,
				pp.operations,
				pp.measure,
				pp.min,
				pp.max,
				pp.period_type,
				pp.period_count
			FROM job_records jr
			LEFT JOIN employees e ON jr.employee = e.id
			LEFT JOIN process_processes_properties ppp ON jr.process = ppp.process_id
			LEFT JOIN process_properties pp ON ppp.property_id = pp.id;
		`)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		defer rows.Close()

		var result []map[string]interface{}

		for rows.Next() {
			var (
				position, client, operations, measure string
				periodType sql.NullString
				min, max sql.NullFloat64
				periodCount sql.NullInt64
			)

			rows.Scan(
				&position, &client,
				&operations, &measure,
				&min, &max, &periodType, &periodCount,
			)

			row := map[string]interface{}{
				"position":   position,
				"client":     client,
				"operations": operations,
				"measure":    measure,
				"min": func() interface{} {
					if min.Valid { return min.Float64 }
						return nil
				}(),
				"max": func() interface{} {
					if max.Valid { return max.Float64 }
						return nil
				}(),
				"period_type": func() interface{} {
					if periodType.Valid { return periodType.String }
				return nil
				}(),
				"period_count": func() interface{} {
					if periodCount.Valid { return periodCount.Int64 }
						return nil
				}(),
			}

			result = append(result, row)
		}

		json.NewEncoder(w).Encode(result)
	}
}