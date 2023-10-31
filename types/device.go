package types

// Device: class defining device info
type Device struct {
	Id           string
	DeviceType   string
	DeviceModel  string
	DeviceSerial string
	DeviceImei   string
	DevicePrice  string
	DeviceDate   string
}

//DeviceSummary : class with device summary
type DeviceSummary struct {
	D                 Device
	DeviceTotal       int
	DeviceAvailable   int
	DeviceAssigned    int
	DeviceUnavailable int
}
