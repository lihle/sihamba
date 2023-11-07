package storage

import (
	"sihmba_server.go/types"
)

// GetAssignLogin: get only employees without logins
func GetAssignLogin() (empoyees []types.Employee, err error) {
	raws, err := db.Query("select e.Employee_id, p.*, e.Employee_position from Employees e " +
		"left join Persons p on e.Person_id = p.Person_id where (select count(*) from Logins l " +
		"where l.Employee_id = e.Employee_id) = 0 and e.Last_date is null")
	if err != nil {
		return
	}
	defer raws.Close()

	for raws.Next() {
		var e types.Employee
		err = raws.Scan(&e.EmployeeID, &e.P.PersonId, &e.P.Firstname, &e.P.Lastname, &e.P.Gender, &e.EmployeePosition)
		if err != nil {
			return
		}
		empoyees = append(empoyees, e)
	}
	return
}
