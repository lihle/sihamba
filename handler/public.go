package handler

import (
	"net/http"

	"sihmba_server.go/html"
	"sihmba_server.go/html/table"
	"sihmba_server.go/storage"
)

// Public : function for public use
func Public(w http.ResponseWriter, r *http.Request) {
	var body string

	body += html.A("/login", "(Admin login)", "right")
	body += html.Br()
	body += html.H2("")
	//
	body += html.Button("/devices", "Sihamba-Sonke Devices")
	body += html.Br()
	body += html.Button("/library", "Sihamba-Sonke Library")
	body += html.Br()
	body += html.Button("/attendances", "Attendance: Programs")
	body += html.Br()
	//
	body += html.H1("")
	//

	view(w, newPage("Welcome", body))
}

//ViewPublicDevices : function for showing quick summary of devices
func ViewPublicDevices(w http.ResponseWriter, r *http.Request) {
	// Declaration of variables
	var body string
	summary, err := storage.GetDeviceSummary()
	if err != nil {
		return
	}
	t := table.New("#", "Devive: Type (example: laptop or tablet etc)", "Device: Total", "Device: Available",
		"Device: Signed-out", "Device: Unavailable")

	//
	// Summary
	body += html.H2("Device: Brief summary")
	for x, s := range summary {
		t.AddRow(x+1, s.D.DeviceType, s.DeviceTotal, s.DeviceAvailable, s.DeviceAssigned, s.DeviceUnavailable)
	}
	body += html.Div(t.HTML("tablesorter"))

	view(w, newPage("Device Summary", body))
}

//ViewPublicLibrary : function for showing quick summary of devices
func ViewPublicLibrary(w http.ResponseWriter, r *http.Request) {
	var body string

	view(w, newPage("Library Summary", body))
}

//ViewPublicAttendances : function for showing quick summary of devices
func ViewPublicAttendances(w http.ResponseWriter, r *http.Request) {
	var body string

	view(w, newPage("Attendance Summary", body))
}
