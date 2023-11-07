package storage

import (
	"sihmba_server.go/types"
)

// GetDevices: returns devices depending on type
func GetDevices(dtype string) (device []types.DeviceLoan, err error) {
	raws, err := db.Query("select d.Device_id, l.Loan_id, d.Device_model, d.Device_serial_no, d.Device_imei_no, "+
		"p.First_name, p.Last_name, s.State_name from Devices d left join Loans l on "+
		"d.Device_id = l.Device_id left join Employees e on l.Employee_id = e.Employee_id left join "+
		"Persons p on e.Person_id = p.Person_id left join States s on s.State_id = l.State_id where d.Device_type = ?", dtype)
	if err != nil {
		return
	}
	defer raws.Close()

	for raws.Next() {
		var d types.DeviceLoan
		err = raws.Scan(&d.D.DeviceId, &d.DeviceLoanId, &d.D.DeviceModel, &d.D.DeviceSerial, &d.D.DeviceImei, &d.E.P.Firstname, &d.E.P.Lastname,
			&d.S.DeviceState)
		if err != nil {
			return
		}
		device = append(device, d)
	}
	return
}

// CheckDeviceLoanout: checks if the device is loaned out and returns 1 or 0
func CheckDeviceLoanout(deviceid string) (value int, err error) {
	raw := db.QueryRow("select count(*) from Loans l left join Devices d on l.Device_id = d.Device_id "+
		"where l.Device_id = ? and l.Loan_in_date is null", deviceid)
	err = raw.Scan(&value)
	return
}
