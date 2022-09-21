package entity

import "time"

type Box struct {
	Id               int       `json:"id,omitempty" db:"id"`
	Code             string    `json:"code,omitempty" db:"code"`
	CreationDate     time.Time `json:"creation_date" db:"creation_date"`
	ContentStateId   int       `json:"content_state_id" db:"content_state_id"`
	UnlimitedStorage bool      `json:"unlimited_storage,required" db:"unlimited_storage"`
	Description      string    `json:"description" db:"description"`
}

func (b *Box) IsValidForSave() bool {
	return len(b.Code) != 0 && b.ContentStateId != 0
}

func (b *Box) IsValidForUpdate() bool {
	return b.Id != 0
}
