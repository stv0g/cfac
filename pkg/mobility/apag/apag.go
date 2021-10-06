package apag

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"log"
)

// curl 'https://www.apag.de/qad_restAPI.php?q=all&d=fulldata&f=json&nocache=862734' \
//   -H 'authority: www.apag.de' \
//   -H 'sec-ch-ua: "Chromium";v="93", " Not;A Brand";v="99"' \
//   -H 'accept: application/json, text/javascript, */*; q=0.01' \
//   -H 'x-requested-with: XMLHttpRequest' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'user-agent: Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://www.apag.de/' \
//   -H 'accept-language: en-US,en;q=0.9' \
//   -H 'cookie: has_js=1' \
//   --compressed
//
// Response:
// [
//    {
//       "ident":"1584",
//       "date":"2021-09-17 18:27:01",
//       "max":"1240",
//       "percent":"21.61",
//       "count":"972",
//       "free":"268",
//       "full":"0",
//       "trend":"up"
//    },
//    {
//       "ident":"1703",
//       "date":"2021-09-17 18:27:01",
//       "max":"999",
//       "percent":"66.47",
//       "count":"335",
//       "free":"664",
//       "full":"0",
//       "trend":"up"
//    },
//    {
//       "ident":"1704",
//       "date":"2021-09-17 18:27:01",
//       "max":"600",

// var houses = {
//  "house_110":{
// 	"nid":"110",
// 	"ident":"6950",
// 	"url":"\/parkobjekte\/parkdeck-stadtgalerie",
// 	"title":"Parkdeck StadtGalerie",
// 	"lat":"51.6524610000",
// 	"lon":"7.3416480000",
// 	"capacity":"207",
// 	"free":"200",
// 	"trend":"eq",
// 	"trend_html":"<span class=\"free\">200<span class=\"trend-eq\">Tendenz: gleichbleibend<\/span><\/span>",
// 	"chartxs_url":"https:\/\/www.apag.de\/qad_restAPI.php?q=6950&d=chartxs&max=207",
// 	"open":"24 Std. ge\u00f6ffnet",
// 	"hight":"max. 2,30 m Einfahrtsh\u00f6he",
// 	"type":"Parkdeck"
//  }
// }

const (
	Url        = "https://www.apag.de"
	UrlApi     = Url + "/qad_restAPI.php"
	UrlChartXS = UrlApi + "?q={ident}&d=chartxs&max={max}"
)

type HouseList []House

type House struct {
	NID       uint    `json:"nid,string"`
	Ident     string  `json:"ident"`
	Url       string  `json:"url"`
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat,string"`
	Longitude float32 `json:"lon,string"`
	Capacity  uint    `json:"capacity,string"`
	Open      string  `json:"open"`
	Height    string  `json:"height"`
	Type      string  `json:"type"`

	Stats HouseStats
}

type HouseStats struct {
	Ident   string    `json:"ident"`
	Date    time.Time `json:"date"`
	Max     uint      `json:"max,string"`
	Percent float32   `json:"percent,string"`
	Count   uint      `json:"count,string"`
	Free    uint      `json:"free,string"`
	Full    bool      `json:"full,string"`
	Trend   string    `json:"trend"`
}

func FetchAllHouses() (HouseList, error) {
	resp, err := http.Get(Url)
	if err != nil {
		return nil, err
	}

	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var re = regexp.MustCompile(`(?m)var houses = (.*);$`)

	m := re.FindSubmatch(c)
	if len(m) == 0 {
		return nil, errors.New("failed to find house list")
	}

	j := m[1]

	re = regexp.MustCompile(`(\d+) (\d+)`)
	j = re.ReplaceAll(j, []byte("$1$2"))

	var houseMap map[string]House

	err = json.Unmarshal(j, &houseMap)
	if err != nil {
		return nil, err
	}

	log.Printf("Found houses: %#v", houseMap)

	var houses HouseList = make([]House, len(houseMap))
	for _, h := range houseMap {
		houses = append(houses, h)
	}

	houses.FetchAllStats()

	return houses, nil
}

func (hs HouseList) FetchAllStats() error {
	resp, err := http.Get(UrlApi + "?q=all&d=fulldata&f=json")
	if err != nil {
		return err
	}

	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var allStats []HouseStats

	err = json.Unmarshal(c, &allStats)
	if err != nil {
		return err
	}

	log.Printf("Found house stats: %#v", allStats)

	houseMap := map[string]*House{}
	for _, h := range hs {
		houseMap[h.Ident] = &h
	}

	for _, s := range allStats {
		h := houseMap[s.Ident]
		h.Stats = s
	}

	return nil
}

func (h *House) FetchStats() error {
	resp, err := http.Get(UrlApi + "?q=" + h.Ident + "&d=fulldata&f=json")
	if err != nil {
		return err
	}

	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(c, &h.Stats)
	if err != nil {
		return err
	}

	return nil
}
