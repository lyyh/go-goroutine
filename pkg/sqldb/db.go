package sqldb

import (
	"database/pkg/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Insert(db *sqlx.DB, p *entity.Person) (int64, error) {
	r, err := db.Exec("insert into person(username,sex,email)values(?,?,?)", p.Username, p.Sex, p.Email)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0, err
	}

	fmt.Println("insert succ:", id)
	return id, nil
}

func Delete(db *sqlx.DB, username string) (int64, error) {
	res, err := db.Exec("delete from person where username=?", username)
	if err != nil {
		fmt.Println("exec failed", err)
		return 0, err
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return 0, err
	}
	fmt.Println("delete succ affected row: ", row)
	return row, nil
}
