package entity

type ContentAction struct {
	Id    int    `json:"id,omitempty" db:"id"`
	Title string `json:"title,omitempty" db:"title"`
}

func (ca *ContentAction) IsValidForSave() bool {
	return ca.Title != ""
}

func (ca *ContentAction) IsValidForUpdate() bool {
	return ca.Id != 0
}
