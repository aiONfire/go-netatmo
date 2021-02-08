package netatmo

// DeviceCollection hold all devices from netatmo account
type DeviceCollection struct {
	Body struct {
		Devices []*Device `json:"devices"`
	}
}

// Device is a station or a module
// ID : Mac address
// StationName : Station name (only for station)
// ModuleName : Module name
// BatteryPercent : Percentage of battery remaining
// WifiStatus : Wifi status per Base station
// RFStatus : Current radio status per module
// Type : Module type :
//  "NAMain" : for the base station
//  "NAModule1" : for the outdoor module
//  "NAModule4" : for the additionnal indoor module
//  "NAModule3" : for the rain gauge module
//  "NAModule2" : for the wind gauge module
// DashboardData : Data collection from device sensors
// DataType : List of available datas
// LinkedModules : Associated modules (only for station)
type Device struct {
	ID             string `json:"_id"`
	StationName    string `json:"station_name"`
	ModuleName     string `json:"module_name"`
	BatteryPercent *int32 `json:"battery_percent,omitempty"`
	WifiStatus     *int32 `json:"wifi_status,omitempty"`
	RFStatus       *int32 `json:"rf_status,omitempty"`
	Type           string
	DashboardData  DashboardData `json:"dashboard_data"`
	//DataType      []string      `json:"data_type"`
	LinkedModules []*Device `json:"modules"`
}
