package netatmo

// DashboardData is used to store sensor values
// Temperature : Last temperature measure @ LastMeasure (in °C)
// Humidity : Last humidity measured @ LastMeasure (in %)
// CO2 : Last Co2 measured @ time_utc (in ppm)
// Noise : Last noise measured @ LastMeasure (in db)
// Pressure : Last Sea level pressure measured @ LastMeasure (in mb)
// AbsolutePressure : Real measured pressure @ LastMeasure (in mb)
// Rain : Last rain measured (in mm)
// Rain1Hour : Amount of rain in last hour
// Rain1Day : Amount of rain today
// WindAngle : Current 5 min average wind direction @ LastMeasure (in °)
// WindStrength : Current 5 min average wind speed @ LastMeasure (in km/h)
// GustAngle : Direction of the last 5 min highest gust wind @ LastMeasure (in °)
// GustStrength : Speed of the last 5 min highest gust wind @ LastMeasure (in km/h)
// LastMeasure : Contains timestamp of last data received
type DashboardData struct {
	Temperature      *float32 `json:"Temperature,omitempty"` // use pointer to detect ommitted field by json mapping
	Humidity         *int32   `json:"Humidity,omitempty"`
	CO2              *int32   `json:"CO2,omitempty"`
	Noise            *int32   `json:"Noise,omitempty"`
	Pressure         *float32 `json:"Pressure,omitempty"`
	AbsolutePressure *float32 `json:"AbsolutePressure,omitempty"`
	Rain             *float32 `json:"Rain,omitempty"`
	Rain1Hour        *float32 `json:"sum_rain_1,omitempty"`
	Rain1Day         *float32 `json:"sum_rain_24,omitempty"`
	WindAngle        *int32   `json:"WindAngle,omitempty"`
	WindStrength     *int32   `json:"WindStrength,omitempty"`
	GustAngle        *int32   `json:"GustAngle,omitempty"`
	GustStrength     *int32   `json:"GustStrength,omitempty"`
	LastMeasure      *int64   `json:"time_utc"`
}
