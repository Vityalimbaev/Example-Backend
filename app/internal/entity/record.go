package entity

type RecordSearchParams struct {
	Id                int    `json:"id,omitempty" db:"id"`
	StartArchivedDate int64  `json:"start_archived_date" db:"start_archived_date"`
	EndArchivedDate   int64  `json:"end_archived_date" db:"end_archived_date"`
	Branch            string `json:"branch,omitempty" db:"branch"`
	StartCreationDate int64  `json:"start_creation_date" db:"start_creation_date"`
	EndCreationDate   int64  `json:"end_creation_date" db:"end_creation_date"`
	Pcode             int64  `json:"pcode,omitempty" db:"pcode"`
	StartLastTreat    int64  `json:"start_last_treat" db:"start_last_treat"`
	EndLastTreat      int64  `json:"end_last_treat" db:"end_last_treat"`
	ContentStateId    int    `json:"content_state_id,omitempty" db:"content_state_id"`
	BoxId             int    `json:"box_id" db:"box_id"`
}

type Record struct {
	Id             int    `json:"id,omitempty" db:"id"`
	ArchivedDate   int64  `json:"archived_date" db:"archived_date"`
	Branch         string `json:"branch,omitempty" db:"branch"`
	CreationDate   int64  `json:"creation_date" db:"creation_date"`
	Pcode          int64  `json:"pcode,omitempty" db:"pcode"`
	LastTreat      int64  `json:"last_treat" db:"last_treat"`
	ContentStateId int    `json:"content_state_id,omitempty" db:"content_state_id"`
	BoxId          int    `json:"box_id" db:"box_id"`
}

func (rec *Record) IsValidForSave() bool {
	return rec.Pcode != 0 && rec.ContentStateId != 0 && rec.BoxId != 0
}

func (rec *Record) IsValidForUpdate() bool {
	return rec.Id != 0
}
