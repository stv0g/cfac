package sensor_community

type SensorDataValues struct {
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
	ID        int64  `json:"id"`
}

type SensorType struct {
	ID           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

type SensorLocation struct {
	ID            int    `json:"id"`
	Indoor        int    `json:"indoor"`
	ExactLocation int    `json:"exact_location"`
	Longitude     string `json:"longitude"`
	Altitude      string `json:"altitude"`
	Latitude      string `json:"latitude"`
	Country       string `json:"country"`
}

type SensorInfo struct {
	Pin        string     `json:"pin"`
	ID         string     `json:"id"`
	SensorType SensorType `json:"sensor_type"`
}

type Sensor struct {
	Sensor       SensorInfo         `json:"sensor"`
	Location     SensorLocation     `json:"location"`
	ID           int64              `json:"id"`
	DataValues   []SensorDataValues `json:"sensordatavalues"`
	Timestamp    string             `json:"timestamp"`
	SamplingRate interface{}        `json:"sampling_rate"`
}
