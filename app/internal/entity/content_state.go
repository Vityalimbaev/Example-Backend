package entity

type ContentState struct {
	Id    int    `json:"id,omitempty" db:"id"`
	Title string `json:"title,omitempty" db:"title"`
}

func (cs *ContentState) IsValidForSave() bool {
	return cs.Title != ""
}

func (cs *ContentState) IsValidForUpdate() bool {
	return cs.Id != 0
}
