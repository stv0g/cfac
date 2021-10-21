package ford_carsharing

// curl 'https://www.ford-carsharing.de/de/rpc' \
//   --data-raw '{"method":"Search.byGeoPosition","params":[{"geoPos":{"radius":"7000","lat":50.790061643,"lng":6.060413354},"maxItems":"200","dateTimeStart":"2021-09-30 12:15:00","dateTimeEnd":"2021-09-30 13:15:00","address":"Aachen, Deutschland","vehicleTypeIds":[],"equipment":[]}],"id":1632996963387}'

const (
	UrlApi = "https://www.ford-carsharing.de/de/rpc"
)
