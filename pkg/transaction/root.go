package transaction

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func TracInsert(db *sqlx.DB) {
	conn, err := db.Begin()
	res, err := conn.Exec("insert into top(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	id, err := res.LastInsertId()
	fmt.Println("LastInsertId", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		// conn.Rollback()
		return
	}
	res, err = conn.Exec("delete from person where user_id=?", 10)
	if err != nil {
		fmt.Println("exec failed, ", err)
		// conn.Rollback()
		return
	}
	conn.Commit()
}
