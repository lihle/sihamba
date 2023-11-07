package handler

import (
	"net/http"

	"sihmba_server.go/html"
	"sihmba_server.go/html/table"
	"sihmba_server.go/storage"
)

// ViewLogin: view page for login
func ViewLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login/")
}

// LoginVerification checks if user has login rights
func LoginVerification(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	userId, err := storage.GetLoginID(username, password)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		count, _ := storage.CheckActiveLogs(userId)
		if count == 0 {
			logId, err := storage.GetLogId(userId)
			if err != nil {
				return
			} else {
				http.Redirect(w, r, "/Admin_page?u="+logId, http.StatusSeeOther)
			}
		} else {
			err = storage.ClearActiveLogs(userId)
			if err != nil {
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

// ViewAdminMenu : view page for admin main page
func ViewAdminMenu(w http.ResponseWriter, r *http.Request) {
	var body string
	u := r.FormValue("u")
	value, err := storage.GetSystemAdmin(u)
	if err != nil {
		return
	}

	//
	body += html.Div(html.A("/logout_func?u="+u, "(logout)"))
	body += html.Br()
	body += viewuserpage(u)

	// Check if user is system admin or not
	if value == 1 {
		body += rootadminoptions(u)
	}

	view(w, newPage("Welcome to User Admin", body))
}

// LogoutFunc: logs user out of the system
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	err := storage.ClearActiveLogs(u)
	if err != nil {
		return
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// viewuserpage: returns user page options
func viewuserpage(u string) string {
	// Declaration of variables
	var body string
	devices, _ := storage.GetDeviceSummary()
	attendances, _ := storage.GetAttendanceSummary()
	books, _ := storage.GetBookSummary()

	// tables
	dt := table.New("#", "Devive: Type (example: laptop or tablet etc)", "Device: Total", "Device: Available",
		"Device: Signed-out", "Device: Not working")
	bt := table.New("#", "Book Title", "Book Staorge", "Overall quantity", "Loaned out")
	at := table.New("#", "Program name", "Event/Class name", "Event/Class date", "Event/Class venue",
		"No: of Attendees", "No: of Enrolled")

	// Devices
	body += html.H2("Device summary", "clickable")
	for x, d := range devices {
		dt.AddRow(x+1, html.A("/Admin_page/device?type="+d.D.DeviceType+"&u="+u, d.D.DeviceType), d.DeviceTotal,
			d.DeviceAvailable, d.DeviceAssigned, d.DeviceUnavailable)
	}
	body += html.Div(dt.HTML("tablesorter"), "hidden")
	body += html.Br()

	// Attendances
	body += html.H2("Attendance summary", "clickable")
	for y, a := range attendances {
		at.AddRow(y+1, a.Ev.Pr.ProgramName, a.Ev.EventName, html.A("/Admin_page/event_?eid="+a.Ev.EventId+"&u="+u,
			a.Ev.EventDate), a.Ev.EventVenue, a.NumberAttended, a.NumberEnrolled)
	}
	body += html.Div(at.HTML("tablesorter")+html.Button("/Admin_page/add_event?u="+u, "Add New Event/Class"), "hidden")
	body += html.Br()

	// Books
	body += html.H2("Book summary", "clickable")
	for z, b := range books {
		bt.AddRow(z+1, html.A("/Admin_page/book_?type="+b.B.BookTitle+"&u="+u, b.B.BookTitle), b.B.BookStorage, b.BookTotal,
			b.BookAssigned)
	}
	body += html.Div(bt.HTML("tablesorter"), "hidden")
	body += html.Br()

	return body
}

// rootadminoptions: adds more options for the root admin
func rootadminoptions(u string) string {
	var body string
	body += html.H2("System admin")
	body += html.Button("/Admin_page/add_login?u="+u, "Add New Login")
	body += html.Br()
	body += html.Button("/Admin_page/add_program?u="+u, "Add New Program")
	body += html.Br()
	body += html.Button("/Admin_page/add_term?u="+u, "Add New Term")
	body += html.Br()
	body += html.Button("/Admin_page/add_states?u="+u, "Add Device State")
	body += html.Br()
	body += html.Button("/Admin_page/add_condition?u="+u, "Add Book Condition")
	body += html.Br()
	return body
}
