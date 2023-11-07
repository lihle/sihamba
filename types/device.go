package types

// Device: class defining device info
type Device struct {
	DeviceId     string
	DeviceType   string
	DeviceModel  string
	DeviceSerial string
	DeviceImei   string
	DevicePrice  string
	DeviceDate   string
}

// DeviceSummary : class with device summary
type DeviceSummary struct {
	D                 Device
	DeviceTotal       int
	DeviceAvailable   int
	DeviceAssigned    int
	DeviceUnavailable int
}

// DeviceStates: class
type DeviceStates struct {
	DeviceStateId string
	DeviceState   string
}

// DeviceLoan: class
type DeviceLoan struct {
	DeviceLoanId string
	D            Device
	E            Employee
	S            DeviceStates
	LoanoutDate  string
	LoaninDate   string
}
