package models

type User struct {
	ID    uint
	Name  string
	Pass  []byte
	Admin bool
}

// UserExternal Model
//
// The User holds information about permission and other stuff.
type UserExternal struct {
	// The user id.
	ID uint `json:"id"`
	// The user name. for login.
	Name string `binding:"required" json:"name" query:"name" form:"name"`
	// If the user is an administrator.
	Admin bool `json:"admin" form:"admin" query:"admin"`
}
