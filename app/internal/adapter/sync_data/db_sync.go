package sync_data

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Content struct {
	BoxId         string    `db:"box_id"`
	Code          string    `db:"code"`
	ArchivedDate  time.Time `db:"archived_date"`
	Branch        string    `db:"branch"`
	CreationDate  time.Time `db:"creation_date"`
	LastTreat     time.Time `db:"last_treat"`
	Pcode         int64     `db:"pcode"`
	SyncLastTreat string    `db:"sync_last_treat"`
}

type Box struct {
	Code         string    `db:"code"`
	CreationDate time.Time `db:"creation_date"`
	Description  string    `db:"description"`
	IsUnlimited  bool      `db:"is_unlimited_storage"`
}

type User struct {
	Name          string `db:"name"`
	Email         string `db:"email"`
	Username      string `db:"username"`
	DefaultBranch string `db:"default_branch"`
}

//func SyncDb() {
//	oldDB, err := sqlx.Connect("mysql", "app-base:5zZcCfvcnuK3@(localhost:3306)/archive-old-prod?parseTime=true")
//	if err != nil {
//		log.Fatalln(err)
//	}

//newDB := database.GetDbConnection()

//	syncUser(oldDB, newDB)
//syncBoxDB(oldDB, newDB)
//syncRecord(oldDB, newDB)
//}
//
//func syncUser(oldDB, newDB *sqlx.DB) {
//	userrepo := user.NewRepository(newDB)
//
//	query := "SELECT default_branch, email, name,  username FROM user WHERE user.username != 'admin' "
//
//	var users []User
//	err := oldDB.Select(&users, query)
//
//	fmt.Println(err)
//
//	for _, user := range users {
//		_, err = userrepo.InsertUser(&entity.User{
//			Name:         user.Name,
//			Username:     user.Username,
//			Email:        user.Email,
//			Password:     "",
//			Branch:       user.DefaultBranch,
//			RoleId:       2,
//			ActiveStatus: false,
//		})
//		fmt.Println(err)
//	}
//}
//
//func syncRecord(oldDB, newDB *sqlx.DB) {
//	query := `SELECT box_id, code, archived_date, branch, content.creation_date,last_treat, pcode, COALESCE(sync_last_treat, '') as sync_last_treat
//				FROM content LEFT JOIN box ON content.box_id = box.id`
//
//	var res []Content
//	err := oldDB.Select(&res, query)
//
//	fmt.Println(err)
//
//	for i, e := range res {
//		query = "INSERT INTO record (archived_date, branch, creation_date,last_treat, pcode, content_state_id, box_id) VALUES ($1, $2, $3,$4, $5, $6, $7 )"
//		_, err = newDB.Exec(query, e.ArchivedDate, e.Branch, e.CreationDate, e.LastTreat, e.Pcode, 1, e.BoxId)
//		fmt.Println(err)
//
//		if i == 200 {
//			break
//		}
//	}
//
//}
//
//func syncBoxDB(oldDB, newDB *sqlx.DB) {
//	query := "SELECT code, creation_date, IF(is_unlimited_storage = 1, true, false)  as is_unlimited_storage, COALESCE(description,'') as description FROM box"
//
//	var res []Box
//	err := oldDB.Select(&res, query)
//	fmt.Println(err)
//
//	for _, e := range res {
//		query = "INSERT INTO box (code,description, is_unlimited_storage, content_state_id) VALUES ($1, $2, $3, 1)"
//		_, err = newDB.Exec(query, e.Code, e.Description, e.IsUnlimited)
//		fmt.Println(err)
//	}
//
//}
