package odlinfo

// curl 'https://odlinfo.bfs.de/json2d69b40fc043fb302fc327c5424ec205/stamm.json' \
//   -H 'sec-ch-ua: "Chromium";v="93", " Not;A Brand";v="99"' \
//   -H 'Accept: application/json, text/javascript, */*; q=0.01' \
//   -H 'Referer: https://odlinfo.bfs.de/DE/aktuelles/messstelle/053130003.html' \
//   -H 'X-Requested-With: XMLHttpRequest' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'User-Agent: Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   --compressed

// https://odlinfo.bfs.de/DE/aktuelles/messstelle/053130003.html

type StammInfo struct {
	Ort    string  `json:"ort"`
	Kenn   string  `json:"kenn"`
	Plz    string  `json:"plz"`
	Status int     `json:"status"`
	Kid    int     `json:"kid"`
	Hoehe  int     `json:"hoehe"`
	Lon    float64 `json:"lon"`
	Lat    float64 `json:"lat"`
	Mw     float64 `json:"mw"`
}

type StammDaten struct {
	Stamm StammInfo `json:"stamm"`
	Mw1H  struct {
		T  []string      `json:"t"`
		Mw []float64     `json:"mw"`
		Ps []int         `json:"ps"`
		Tr []string      `json:"tr"`
		R  []interface{} `json:"r"`
	} `json:"mw1h"`
	Mw24H struct {
		T  []string  `json:"t"`
		Mw []float64 `json:"mw"`
	} `json:"mw24h"`
}
