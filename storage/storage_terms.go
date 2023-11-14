package storage

import (
	"fmt"
)

// InsertTerm: insert a new term
func InsertTerm(number int, start, end string) (string, error) {
	res, err := db.Exec("insert into Terms set Term_number = ?, Term_start_date = ?, Term_end_date = ?", number, start, end)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	return fmt.Sprint(id), err
}
