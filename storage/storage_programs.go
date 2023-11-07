package storage

import (
	"fmt"
)

// InsertProgram: insert new program entry
func InsertProgram(name, brief string) (string, error) {
	res, err := db.Exec("insert into Programs set Program_name = ?, Program_about = ?", name, brief)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	return fmt.Sprint(id), err
}
