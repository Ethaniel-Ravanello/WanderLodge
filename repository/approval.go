package repository

import "database/sql"

func CreateApproval(db *sql.DB, listingId int, approvalType string, approverId sql.NullInt32, status string) (err error) {
	sql := `INSERT INTO approvals ("approvalTypeId", "approvalType", "approverId", "status", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id;`

	_, err = db.Exec(sql, listingId, approvalType, approverId, status)

	return err
}

func ActionApproval(db *sql.DB, approvalId int, newStatus string) (err error) {

	sqlUpdateApproval := `UPDATE approvals SET "status" = $1, "updatedAt" = NOW() WHERE "id" = $2`
	_, err = db.Exec(sqlUpdateApproval, newStatus, approvalId)
	if err != nil {
		return err
	}

	var approvalType string
	var relatedId int

	sqlSelect := `SELECT "approvalType", "approvalTypeId" FROM approvals WHERE "id" = $1`
	err = db.QueryRow(sqlSelect, approvalId).Scan(&approvalType, &relatedId)
	if err != nil {
		return err
	}

	var sqlUpdateRelated string
	if approvalType == "listing" {
		sqlUpdateRelated = `UPDATE listings SET "approvalStatus" = $1 WHERE "id" = $2`
	} else if approvalType == "booking" {
		sqlUpdateRelated = `UPDATE bookings SET "status" = $1 WHERE "id" = $2`
	} else {
		return nil
	}

	_, err = db.Exec(sqlUpdateRelated, newStatus, relatedId)
	return err
}
