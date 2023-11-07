package handler

import (
	"html/template"
	"net/http"

	"sihmba_server.go/html"
)

// ViewAddTerm: views add new term
func ViewAddTerm(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	var form string
	nums := []string{"1", "2", "3", "4"}

	//
	if set(r.FormValue("submit")) && !set(r.FormValue("error")) {

		http.Redirect(w, r, "/Admin_page?u="+u, http.StatusSeeOther)
	}

	form += html.Div(html.Button("/Admin_page?u="+u, "Admin home page"), "right")
	form += html.Br()
	form += html.H2("Add a new term")
	form += html.Div(html.LabelSelect("Term number: ", "number", nums, nums))
	form += html.Div(html.LabelDate("Term start date: ", "startdate"))
	form += html.Div(html.LabelDate("Term end date: ", "enddate"))

	page := multiPartForm("Add new program", form)

	page.Error = template.HTML(r.FormValue("error"))
	view(w, page)
}
