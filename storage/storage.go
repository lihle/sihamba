package storage

import (
	"database/sql"

	"sihmba_server.go/types"
	//this driver is needed to run sql queries to mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Init connects to the database and returns an error upon failure
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

// GetBookSummary: function returns books summary for all Books
func GetBookSummary() (Books []types.BookSummary, err error) {
	raws, err := db.Query("select distinct(b.Book_title), b.Storage_reference, " +
		"(select count(*) from Books n where n.Book_title = b.Book_title), " +
		"(select count(*) from Books c left join B_loans t on t.Book_id = c.Book_id " +
		"where c.Book_title = b.Book_title and t.Loan_in_date is null and t.Loan_out_date is not null) " +
		"from Books b")
	if err != nil {
		return
	}
	defer raws.Close()

	for raws.Next() {
		var b types.BookSummary
		err = raws.Scan(&b.B.BookTitle, &b.B.BookStorage, &b.BookTotal, &b.BookAssigned)
		if err != nil {
			return
		}
		Books = append(Books, b)
	}
	return
}

// GetAttendanceSummary: function returns attendences
func GetAttendanceSummary() (attendances []types.Attendance, err error) {
	raws, err := db.Query("select e.Event_id, p.Program_name, e.Event_name, e.Event_date, e.Event_venue, " +
		"(select count(*) from Attendances a where a.Event_id = e.Event_id and a.Attended = '1') as 'Attended', " +
		"(select count(*) from Enrollments j where j.Program_id = p.Program_id) as 'Enrolled' " +
		"from Events e left join Programs p on e.Program_id = p.Program_id order by e.Event_date, p.Program_name")
	if err != nil {
		return
	}
	defer raws.Close()

	for raws.Next() {
		var a types.Attendance
		err = raws.Scan(&a.Ev.EventId, &a.Ev.Pr.ProgramName, &a.Ev.EventName, &a.Ev.EventDate, &a.Ev.EventVenue, &a.NumberAttended, &a.NumberEnrolled)
		if err != nil {
			return
		}
		attendances = append(attendances, a)
	}
	return
}
