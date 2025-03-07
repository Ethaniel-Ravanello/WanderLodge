package repository

import (
	"database/sql"
	"wanderloge/structs"
)

func AddBookListing(db *sql.DB, booking structs.Booking) (err error) {
	sql := `INSERT INTO bookings ("guestId", "listingId", "persons", "startDate", "endDate", "status") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	_, err = db.Exec(sql, booking.GuestId, booking.ListingId, booking.Persons, booking.StartDate, booking.EndDate, "pending")

	return err
}

func GetAllBookingForHost(db *sql.DB, hostId int) (bookings []structs.Booking, err error) {
	sql := `SELECT * FROM bookings WHERE hostId=$1`

	rows, err := db.Query(sql, hostId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempListing structs.Booking
		err = rows.Scan(&tempListing)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, tempListing)
	}
	return bookings, nil
}

func GetAllBookingForGuest(db *sql.DB, guestId int) (bookings []structs.Booking, err error) {
	sql := `SELECT * FROM bookings WHERE guestId=$1`

	rows, err := db.Query(sql, guestId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempListing structs.Booking
		err = rows.Scan(&tempListing)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, tempListing)
	}
	return bookings, nil
}

func GetBookingDetailById(db *sql.DB, id int) (bookings []structs.Booking, err error) {
	sql := `SELECT FROM bookings WHERE id=$1`

	err = db.QueryRow(sql, id).Scan(&bookings)

	if err != nil {
		return nil, err
	}

	return bookings, nil
}
