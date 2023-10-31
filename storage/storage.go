package storage

import (
	"database/sql"

	"sihmba_server.go/types"
	//this driver is needed to run sql queries to mysql
)

var db *sql.DB

//Init connects to the database and returns an error upon failure
func Init() error {
	var err error
	db, err = conn()
	return err
}

func conn() (*sql.DB, error) {
	return sql.Open("mysql", "root:lihle@(localhost:3306)/sonke?charset=utf8")
}

func valueOrWildcard(in string) string {
	if in == "" {
		return "%"
	}
	return in
}

// GetDeviceSummary: function returns device summary for all devices
func GetDeviceSummary() (Summary []types.DeviceSummary, err error) {
	raws, err := db.Query("select distinct(d.Device_type), (select count(*) from Devices t " +
		"where t.Device_type = d.Device_type), (select count(*) from Devices a left join Loans l " +
		"on a.Device_id = l.Device_id where a.Device_type = d.Device_type and l.Loan_in_date is null), " +
		"(select count(*) from Devices a left join Loans l on a.Device_id = l.Device_id " +
		"where a.Device_type = d.Device_type and l.Loan_in_date is not null), (select count(*) " +
		"from Devices u left join Loans o on o.Device_id = u.Device_id left join States s " +
		"on s.State_id = o.State_id where s.State_name like 'Not%' ) from Devices d;")
	if err != nil {
		return
	}
	defer raws.Close()

	for raws.Next() {
		var s types.DeviceSummary
		err = raws.Scan(&s.D.DeviceType, &s.DeviceTotal, &s.DeviceAvailable, &s.DeviceAssigned, &s.DeviceUnavailable)
		if err != nil {
			return
		}
		Summary = append(Summary, s)
	}
	return
}
