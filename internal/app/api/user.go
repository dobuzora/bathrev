package api

import (
	"github.com/dobuzora/bathrev/internal/app/models"
	"net/http"
)

type UserDatabase interface {
	GetUsers() []*models.User
	GetUserByID(id uint) *models.User
	GetUserByName(name string) *models.User
	DeleteUserByID(id uint) *models.User
	UpdateUser(user *models.User)
	CreateUser(user *models.User) error
	CountUser(condition ...interface{}) int
}

type UserChangeNotifier struct {
	userDeletedCallbacks []func(uid uint) error
	userAddedCallbacks   []func(uid uint) error
}

type UserAPI struct {
	DB                 UserDatabase
	PasswordStrength   int
	UserChangeNotifier *UserChangeNotifier
}

func (a *UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := a.DB.GetUsers()

	var resp []*models.UserExternal
	for _, user := range users {
		resp = append(resp, toExternalUer(user))
	}

}

func toExternalUer(internal *models.User) *models.UserExternal {
	return &models.UserExternal{
		Name:  internal.Name,
		Admin: internal.Admin,
		ID:    internal.ID,
	}
}
