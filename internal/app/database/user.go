package database

import "github.com/dobuzora/bathrev/internal/app/models"

// GetUser returns all users.
func (d *GormDatabase) GetUsers() []*models.User {
	var users []*models.User
	d.DB.Find(&users)
	return users
}
