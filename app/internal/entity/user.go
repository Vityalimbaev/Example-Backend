package entity

type LoginForm struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type UserSession struct {
	Id           int    `json:"-" db:"id"`
	RefreshToken string `json:"-" db:"refresh_token"`
	DateExpire   int64  `json:"-" db:"date_expire"`
	UserId       int    `json:"-" db:"account_id"`
}

type UserSearchParams struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username,omitempty" db:"username"`
	Email    string `json:"email,omitempty" db:"email"`
	Branch   string `json:"branch" db:"branch"`
	RoleId   int    `json:"role_id" db:"role_id"`
}

type User struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Username     string `json:"username,omitempty" db:"username"`
	Email        string `json:"email,omitempty" db:"email"`
	Password     string `json:"password,omitempty" db:"password"`
	Branch       string `json:"branch" db:"branch"`
	RoleId       int    `json:"role_id" db:"role_id"`
	RoleTitle    string `json:"role_title,omitempty" db:"role_title"`
	CreationDate int64  `json:"creation_date" db:"creation_date"`
	ActiveStatus bool   `json:"active_status" db:"active_status"`
}

func (u *User) IsValidForSave() bool {
	return u.Name != "" && u.Username != "" && u.Email != "" && len(u.Password) > 4 && u.RoleId != 0
}

func (u *User) IsValidForUpdate() bool {
	return u.Id != 0
}
