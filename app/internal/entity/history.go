package entity

type ActionHistory struct {
	Id              int    `json:"id" db:"id"`
	ContentActionId int    `json:"content_action_id" db:"content_action_id"`
	BoxId           int    `json:"box_id" db:"box_id"`
	RecordId        int    `json:"record_id" db:"record_id"`
	Datetime        int64  `json:"datetime" db:"datetime"`
	Description     string `json:"description" db:"description"`
	UserId          int    `json:"user_id" db:"user_id"`
}

func (ah *ActionHistory) IsValidForSave() bool {
	return ah.ContentActionId > 0 && (ah.BoxId != 0 || ah.RecordId != 0) && ah.UserId > 0
}
