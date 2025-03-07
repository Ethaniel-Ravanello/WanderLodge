package repository

import "database/sql"

func CreateApproval(db *sql.DB, listingId int, approvalType string, approverId sql.NullInt32, status string, createdAt interface{}, updatedAt interface{}) (err error) {
	sql := `INSERT INTO approvals ("approvalTypeId", "approvalType", "approverId", "status", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id;`

	_, err = db.Exec(sql, listingId, approvalType, approverId, status)

	return err
}
