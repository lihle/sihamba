package storage

import (
	"fmt"
)

// GetLoginID: returns loginID from username & password
func GetLoginID(username, password string) (id string, err error) {
	raw := db.QueryRow("select l.Login_id from Logins l where l.Username like ? and l.Password like ?", username, password)
	err = raw.Scan(&id)
	return
}

// GetLogId : insert a new log when logging in and returns Log ID
func GetLogId(loginId string) (string, error) {
	res, err := db.Exec("insert into Logs set Login_id = ?, Login_date = now()", loginId)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	return fmt.Sprint(id), err
}

// CheckActiveLogs: returns 0 if there are no active logs otherwise
func CheckActiveLogs(userId string) (count int, err error) {
	raw := db.QueryRow("select count(*) from Logs l where l.Login_id = ? and l.Login_status like 'Active'", userId)
	err = raw.Scan(&count)
	return
}

// ClearActiveLogs: clears all active logs for user id
func ClearActiveLogs(userId string) error {
	_, err := db.Exec("Update Logs set Login_status = 'Inactive', Logout_date = now() where Login_id = ?", userId)
	return err
}

// GetSystemAdmin: return 1 for System admin & 0 for User admin
func GetSystemAdmin(logId string) (admin int, err error) {
	raw := db.QueryRow("select l.Admin from Logs lo left join Logins l on lo.Login_id = l.Login_id where lo.Log_id = ?", logId)
	err = raw.Scan(&admin)
	return
}

// InsertLogin: inserts a new login
func InsertLogin(employeeid, name string, isadmin int) (string, error) {
	res, err := db.Exec("insert into Logins set Employee_id = ?, Username = ?, Password = ?, Admin = ?", employeeid, name, name+"@123", isadmin)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	return fmt.Sprint(id), err
}
