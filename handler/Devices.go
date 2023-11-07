package handler

import (
	"net/http"

	"sihmba_server.go/html"
	"sihmba_server.go/html/table"
	"sihmba_server.go/storage"
)

// ViewDevices: view devices depending on type
func ViewDevices(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	dtype := r.FormValue("type")
	var body string
	devices, err := storage.GetDevices(dtype)
	if err != nil {
		return
	}
	t := table.New("#", "Device model", "Serial no.", "Imei number", "Full-name", "Device state")

	//
	body += html.Div(html.Button("/Admin_page?u="+u, "Home Page", "green"), "right")
	body += html.Div(html.Button("/logout_func?u="+u, "Log-out", "yellow"), "left")

	// List
	body += html.H2("Breakdown of " + dtype)
	for x, d := range devices {
		value, _ := storage.CheckDeviceLoanout(d.D.DeviceId)

		if value == 0 {
			t.AddRow(x+1, html.A("/Admin_page/device/"+d.D.DeviceId+"?u="+u, d.D.DeviceModel), d.D.DeviceSerial,
				d.D.DeviceImei, d.E.P.Firstname+" "+d.E.P.Lastname, d.S.DeviceState)
		} else {
			t.AddRow(x+1, d.D.DeviceModel, d.D.DeviceSerial, d.D.DeviceImei,
				html.A("/Admin_page/device_loan/"+d.DeviceLoanId+"?u="+u, d.E.P.Firstname+" "+d.E.P.Lastname), d.S.DeviceState)
		}

	}
	body += html.Div(t.HTML("tablesorter"))

	view(w, newPage("All "+dtype, body))
}

//
