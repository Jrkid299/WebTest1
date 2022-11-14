package data

import (
	"database/sql"

	"AWtest1.jalenlamb.net/internals/validator"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"age"`
	Email    string `json:"email"`
}

func ValidateUser(v *validator.Validator, user *User) {
	// Use the Check() method to execute our validation checks
	v.Check(user.Username != "", "name", "must be provided")
	v.Check(len(user.Username) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) <= 200, "email", "must not be more than 200 bytes long")

}

// Define a SchoolModel which wraps a sql.DB connection pool
type UserModel struct {
	DB *sql.DB
}

// Insert() allows us  to create a new School
func (m UserModel) Insert(user *User) error {
	query := `
		INSERT INTO schools (username, email)
		VALUES ($1, $2)
		RETURNING id, 
	`
	// Collect the data fields into a slice
	args := []interface{}{
		user.Username, user.Email,
	}
	return m.DB.QueryRow(query, args...).Scan(&user.Id)
}

// Get() allows us to retrieve a specific School
func (m UserModel) Get(id int64) (*User, error) {
	return nil, nil
}

// Update() allows us to edit/alter a specific School
func (m UserModel) Update(user *User) error {
	return nil
}

// Delete() removes a specific School
func (m UserModel) Delete(id int64) error {
	return nil
}
