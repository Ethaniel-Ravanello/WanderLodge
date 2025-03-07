package repository

import (
	"database/sql"
	"wanderloge/structs"
)

func SignUp(db *sql.DB, users structs.User) (int, error) {
	sql := `INSERT INTO users ("firstName", "lastName", "email", "phoneNumber", "roles", "password") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	var listingId int
	err := db.QueryRow(sql, users.FirstName, users.LastName, users.Email, users.PhoneNumber, users.Roles, users.Password).Scan(&listingId)

	if err != nil {
		panic(err)
	}

	return listingId, nil
}

func SignIn(db *sql.DB, userName string) string {
	var hashPass string
	sql := `SELECT "password" FROM users WHERE "firstName" = $1`

	errs := db.QueryRow(sql, userName).Scan(&hashPass)
	if errs != nil {
		return "Error Getting User Data"
	}

	return hashPass
}

func GetUsers(db *sql.DB) (users []structs.User, err error) {
	sql := `SELECT * FROM users`

	rows, errs := db.Query(sql)
	if errs != nil {
		return nil, errs
	}
	defer rows.Close()

	for rows.Next() {
		var tempUser structs.User

		err = rows.Scan(&tempUser.Id, &tempUser.FirstName, &tempUser.LastName, &tempUser.Email,
			&tempUser.PhoneNumber, &tempUser.Roles, &tempUser.Password)

		if err != nil {
			return nil, err
		}
		users = append(users, tempUser)
	}
	return users, nil
}

func GetUserById(db *sql.DB, id int, firstName string) (user structs.User, err error) {
	sql := `SELECT * FROM users WHERE id = $1 OR "firstName" = $2;
`

	errs := db.QueryRow(sql, id, firstName).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Roles, &user.Password)
	if errs != nil {
		return structs.User{}, errs
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user structs.User, userId int) (err error) {
	var tempRole string
	sqlUpdateUser := `UPDATE users set "firstName" = $1, "lastName" = $2, "email" = $3, "phoneNumber" = $4, "roles" = $5 WHERE "id" = $6`
	sqlGetRole := `SELECT "roles" from users WHERE "id" = $1`

	errGetRole := db.QueryRow(sqlGetRole, userId).Scan(&tempRole)
	if errGetRole != nil {
		return errGetRole
	}

	_, errs := db.Exec(sqlUpdateUser, user.FirstName, user.LastName, user.Email, user.PhoneNumber, tempRole, userId)

	return errs
}

func DeleteUser(db *sql.DB, userId int) (err error) {
	sql := `DELETE FROM users WHERE id = $1`

	_, errs := db.Exec(sql, userId)
	if errs != nil {
		return err
	}
	return nil
}
