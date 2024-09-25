package models

type UserRole string

const (
	UserRole_User  UserRole = "user"
	UserRole_Admin UserRole = "admin"
)

type UserModel struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
	Email     string   `json:"email"`
	CreatedAt int64    `json:"createdAt"`
}
