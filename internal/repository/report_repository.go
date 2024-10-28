package repository

import (
	"database/sql"

	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/shared/custom"
)

type ReportRepository interface {
	List(userId, startDate, endDate string) ([]custom.ReportResp, error)
}

type reportRepository struct {
	db  *sql.DB
	log *logger.Logger
}

func (r *reportRepository) List(userId, startDate, endDate string) ([]custom.ReportResp, error) {
	selectQuery := `
		SELECT
			p.name_provider,
			COUNT(t.transaction_id)
		FROM transactions t
		JOIN mst_user u ON t.id_user = u.id_user
		JOIN mst_merchant m ON t.id_merchant = m.id_merchant
		JOIN transaction_detail td ON t.transaction_id = td.transaction_id
		JOIN mst_product p ON td.id_product = p.id_product
		WHERE m.id_merchant = (
			SELECT
				m.id_merchant
			FROM mst_merchant m
			WHERE m.id_user = $1
		)
		AND t.transaction_date >= $2
		AND t.transaction_date <= $3
		GROUP BY p.name_provider
		ORDER BY 2 DESC;`

	r.log.Info("Starting to retrive report of all transactions in the repository layer", nil)

	rows, err := r.db.Query(selectQuery, userId, startDate, endDate)
	if err != nil {
		r.log.Error("Failed to retrieve the report of transactions", err)
		return nil, err
	}
	defer rows.Close()

	var reportSlice []custom.ReportResp

	for rows.Next() {
		var report custom.ReportResp
		if err := rows.Scan(
			&report.ProviderName,
			&report.Count,
		); err != nil {
			r.log.Error("Failed to scan report of transactions", err)
			return nil, err
		}
		reportSlice = append(reportSlice, report)
	}

	if err := rows.Err(); err != nil {
		r.log.Error("Failed to scan report of transactions", err)
		return nil, err
	}

	return reportSlice, nil
}

func NewReportRepository(db *sql.DB, log *logger.Logger) ReportRepository {
	return &reportRepository{db: db, log: log}
}
