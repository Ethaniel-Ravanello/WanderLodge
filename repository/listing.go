package repository

import (
	"database/sql"
	"fmt"
	"wanderloge/structs"
)

func CreateListing(db *sql.DB, listing structs.Listing) (int, string, string, error) {
	sqlStatement := `
		INSERT INTO listings ("hostid", "title", "description", "location", "address", "maxPeople", "pricePerNight", "createdAt", "approvalStatus")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), $8)
		RETURNING id, "createdAt", "approvalStatus";`

	var listingID int
	var createdAt string
	var approvedStatus string
	err := db.QueryRow(sqlStatement,
		listing.HostId, listing.Title, listing.Description, listing.Location,
		listing.Address, listing.MaxPeople, listing.PricePerNight, "pending").Scan(&listingID, &createdAt, &approvedStatus)

	if err != nil {
		return 0, "", "", fmt.Errorf("error creating listing: %w", err)
	}

	return listingID, createdAt, approvedStatus, nil
}

func GetListings(db *sql.DB, status string, id int) ([]structs.Listing, error) {
	var listings []structs.Listing
	sql := `SELECT * FROM listings WHERE "approvalStatus"=$1 AND id=$2`

	rows, err := db.Query(sql, status, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tempListing structs.Listing
		err = rows.Scan(&tempListing.Id, &tempListing.HostId, &tempListing.Title, &tempListing.Description, &tempListing.Location, &tempListing.Address, &tempListing.MaxPeople, &tempListing.PricePerNight, &tempListing.CreatedAt, &tempListing.ApprovalStatus)
		if err != nil {
			return nil, err
		}
		listings = append(listings, tempListing)
	}
	return listings, nil
}

func GetListingByListingId(db *sql.DB, listingId int) (listings structs.Listing, err error) {
	sql := `SELECT * FROM listings WHERE id = $1`

	err = db.QueryRow(sql, listingId).Scan(&listings.Id, &listings.HostId, &listings.Title, &listings.Description, &listings.Location, &listings.Address, &listings.MaxPeople, &listings.PricePerNight, &listings.CreatedAt, &listings.ApprovalStatus)

	if err != nil {
		return listings, err
	}
	return listings, nil
}

func GetListingByHostId(db *sql.DB, hostId int) (listings structs.Listing, err error) {
	sql := `SELECT * FROM listings WHERE hostId = $1`

	err = db.QueryRow(sql, hostId).Scan(&listings.Id, &listings.HostId, &listings.Title, &listings.Description, &listings.Location, &listings.Address, &listings.MaxPeople, &listings.PricePerNight, &listings.CreatedAt, &listings.ApprovalStatus)

	if err != nil {
		return listings, err
	}

	return listings, nil
}

func UpdateListing(db *sql.DB, listing structs.Listing, listingId int) (err error) {
	sql := `UPDATE listings SET 
                    title=$1, 
                    description=$2,
                    location=$3,
                    address=$4,
                    "maxPeople"=$5,
                    "pricePerNight"=$6
			WHERE id=$7`

	_, err = db.Exec(sql, listing.Title, listing.Description, listing.Location, listing.Address, listing.MaxPeople, listing.PricePerNight, listingId)

	return err
}

func DeleteListing(db *sql.DB, listingId int) (err error) {
	sql := `DELETE FROM listings WHERE id=$1`

	_, err = db.Exec(sql, listingId)

	return err
}
