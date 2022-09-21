package entity

type Role struct {
	Id    int    `json:"id,omitempty" db:"id"`
	Title string `json:"title" db:"title"`
}

func (r *Role) IsValidForSave() bool {
	return r.Title != ""
}

func (r *Role) IsValidForUpdate() bool {
	return r.Id != 0
}
