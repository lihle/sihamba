package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"sihmba_server.go/html"
	"sihmba_server.go/storage"
)

// ViewAddTerm: views add new term
func ViewAddTerm(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	var form string
	nums := []string{"1", "2", "3", "4"}

	//
	if set(r.FormValue("submit")) && !set(r.FormValue("error")) {
		number := r.FormValue("number")
		start := r.FormValue("startdate")
		end := r.FormValue("enddate")

		if set(number, start, end) {
			num, _ := strconv.Atoi(number)
			_, err := storage.InsertTerm(num, start, end)
			if err != nil {
				return
			}
			http.Redirect(w, r, "/Admin_page?u="+u, http.StatusSeeOther)
		}
		form += html.Error("Information missing")
	}

	form += html.Div(html.Button("/Admin_page?u="+u, "Admin home page"), "right")
	form += html.Br()
	form += html.H2("Add a new term")
	form += html.Div(html.LabelSelect("Term number: ", "number", nums, nums))
	form += html.Div(html.LabelDate("Term start date: ", "startdate"))
	form += html.Div(html.LabelDate("Term end date: ", "enddate"))

	page := multiPartForm("Add new Term", form)

	page.Error = template.HTML(r.FormValue("error"))
	view(w, page)
}
