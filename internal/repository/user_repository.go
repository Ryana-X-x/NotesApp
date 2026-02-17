package repository

import (
	"gorm.io/gorm"
	"github.com/aryansehgal-tech/NotesApp/internal/models"
)

/*
UserRepository handles all database operations related to the User model.

Why this exists:
- Keeps database logic separate from business logic (service layer)
- Prevents handlers from directly interacting with GORM
- Makes the code easier to test and maintain
- Follows clean architecture principles

This struct acts as a wrapper around the GORM DB instance.
*/
type UserRepository struct {
	db *gorm.DB // GORM database connection (injected dependency)
}

/*
NewUserRepository is a constructor function.

Why we use a constructor:
- Injects the database dependency
- Ensures repository is initialized properly
- Makes testing easier (can inject mock DB)

Returns:
- Pointer to UserRepository
*/
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

/*
Create inserts a new user record into the database.

Parameters:
- user: pointer to models.User struct (data to be saved)

Flow:
- GORM builds and executes INSERT query
- If insertion fails, GORM sets .Error

Returns:
- nil if successful
- error if something goes wrong

Why we return only error:
- Caller usually already has the user object
- We only need to know if DB operation succeeded
*/
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

/*
FindByEmail retrieves a user by their email address.

Parameters:
- email: string used to search in database

Flow:
1. Create empty user struct
2. Execute SELECT query with WHERE condition
3. Use parameter binding (?) to prevent SQL injection
4. Store result inside user struct

Important:
- If no record is found, GORM returns gorm.ErrRecordNotFound
- Caller must check the returned error before using the user

Returns:
- Pointer to found user
- error (nil if found, otherwise error)
*/
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	// WHERE email = ? ensures safe parameter substitution
	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
