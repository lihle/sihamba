package handler

import (
	"html/template"
	"net/http"
	"strings"

	"sihmba_server.go/html"
	"sihmba_server.go/storage"
	"sihmba_server.go/types"
)

// ViewAssignLogin: views the assign login page
func ViewAssignLogin(w http.ResponseWriter, r *http.Request) {
	var form string
	u := r.FormValue("u")
	employees, _ := storage.GetAssignLogin()
	names := Assignemployees(employees)
	rights := []string{"No", "Yes"}

	//
	if set(r.FormValue("submit")) && !set(r.FormValue("error")) {
		var value int
		word := r.FormValue("name")
		arr := strings.Split(word, ".")
		employeeid := arr[0]
		details := strings.Split(arr[1], " ")
		name := details[0]

		isadmin := r.FormValue("isadmin")

		if isadmin == "Yes" {
			value = 1
		} else {
			value = 0
		}

		//
		_, err := storage.InsertLogin(employeeid, name, value)
		if err != nil {
			return
		} else {
			http.Redirect(w, r, "/Admin_page?u="+u, http.StatusSeeOther)
		}

	}

	//
	form += html.Div(html.Button("/Admin_page?u="+u, "Admin home page"), "right")
	form += html.Br()
	form += html.H2("Assign employee Logins")
	form += html.Br()
	form += html.Div(html.LabelSelect("Employee name: ", "name", names, names))
	form += html.Div(html.LabelSelect("Is System Admin? (Yes: System admin & No: User Admin): ", "isadmin", rights, rights))

	page := multiPartForm("Assign Employee login", form)

	page.Error = template.HTML(r.FormValue("error"))
	view(w, page)
}

// employees
func Assignemployees(employees []types.Employee) (names []string) {
	names = append(names, "")
	for _, e := range employees {
		n := e.EmployeeID + "." + e.P.Firstname + " " + e.P.Lastname
		names = append(names, n)
	}
	return
}
