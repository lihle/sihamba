package handler

import (
	"html/template"
	"net/http"

	"sihmba_server.go/html"
	"sihmba_server.go/storage"
)

// ViewProgram: views add program page
func ViewProgram(w http.ResponseWriter, r *http.Request) {
	var form string
	u := r.FormValue("u")
	//
	if set(r.FormValue("submit")) && !set(r.FormValue("error")) {
		name := r.FormValue("program")
		brief := r.FormValue("description")

		_, err := storage.InsertProgram(name, brief)
		if err != nil {
			return
		} else {
			http.Redirect(w, r, "/Admin_page?u="+u, http.StatusSeeOther)
		}
	}

	form += html.Div(html.Button("/Admin_page?u="+u, "Admin home page"), "right")
	form += html.Br()
	form += html.H2("Add a new program")
	form += html.Div(html.LabelString("Program name: ", "program") + html.Br())
	form += html.Div(html.LabelTextArea("Description (What the program is about): ", "description") + html.Br())

	page := multiPartForm("Add new program", form)

	page.Error = template.HTML(r.FormValue("error"))
	view(w, page)
}
