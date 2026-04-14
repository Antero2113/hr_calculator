package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"project/models"
	"project/utils"
)

func AddRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.Record
		json.NewDecoder(r.Body).Decode(&data)

		department, region := utils.ParseDepartment(data.DepartmentFull)

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		res, err := tx.Exec(
			"INSERT INTO employees (name, position, region, department) VALUES (?, ?, ?, ?)",
			data.Name, data.Position, region, department,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}

		employeeId, _ := res.LastInsertId()

		res, err = tx.Exec("INSERT INTO processes (space_id) VALUES (NULL)")
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}

		processId, _ := res.LastInsertId()

		res, err = tx.Exec(
			"INSERT INTO process_properties (name, operations, measure, min, max, period_type, period_count) VALUES (?, ?, ?, ?, ?, ?, ?)",
			data.ProcessName, data.Operations, data.Measure, data.Min, data.Max, data.PeriodType, data.PeriodCount,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}

		propertyId, _ := res.LastInsertId()

		_, err = tx.Exec(
			"INSERT INTO process_processes_properties (process_id, property_id, value) VALUES (?, ?, '')",
			processId, propertyId,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}

		_, err = tx.Exec(
			"INSERT INTO job_records (employee, client, process) VALUES (?, ?, ?)",
			employeeId, data.Client, processId,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}

		tx.Commit()

		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}