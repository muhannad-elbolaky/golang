package models

import (
	"errors"

	"elbolaky.com/rest-api/db"
	"elbolaky.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save saves the user's information to the database.
//
// It prepares an SQL query to insert the user's email and hashed password into the "users" table.
// If there is an error preparing the query, it returns the error.
// It then executes the query and retrieves the last inserted ID.
// The user's ID is updated with the last inserted ID.
// If there is an error executing the query or retrieving the last inserted ID, it returns the error.
//
// Returns:
// - error: An error if there was a problem saving the user's information.
func (user *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	user.ID = userID
	return err
}

// ValidateCredentials validates the user's credentials by checking if the provided email and password match the records in the database.
//
// It queries the "users" table in the database to retrieve the user's ID and hashed password based on the provided email.
// It then compares the provided password with the retrieved hashed password using the CheckPasswordHash function from the utils package.
// If the passwords match, it returns nil, indicating that the credentials are valid.
// If the passwords do not match or any error occurs during the process, it returns an error indicating the validation failure.
//
// Parameters:
// - None
//
// Returns:
// - error: An error indicating the validation failure, or nil if the credentials are valid.
func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string

	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

// GetAllUsers retrieves all users from the database.
//
// It executes a SELECT query on the "users" table and returns a slice of User structs.
// If there is an error during the query execution, it returns nil and the error.
//
// Returns:
// - []User
func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
