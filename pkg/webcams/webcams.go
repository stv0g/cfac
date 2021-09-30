package webcams

import (
	"net/url"
	"time"
)

// https://www.drehturm-aachen.de/de/livecam/
// https://kullenhof.de/heli/min.php
// https://kullenhof.de/heli/cam_pic.php

// https://www.strassen.nrw.de/de/projekte/a4/autobahnkreuz-aachen-a4-a44-a544/webcams.html
// https://portal1469.webcam-profi.de/

var (
	Webcams []Webcam
)

type WebcamType int

type Webcam struct {
	Type       WebcamType
	Url        *url.URL
	Source     *url.URL
	UpdateRate time.Duration
}

func parseURL(u string) *url.URL {
	i, _ := url.Parse(u)
	return i
}

func init() {
	wetterEdka := parseURL("http://wetter-edka.de/")
	akAachen := parseURL("https://ak-aachen.de/webcams")

	Webcams = []Webcam{
		{
			Url:    parseURL("http://wetter-edka.de/cam1.jpg"),
			Source: wetterEdka,
		},
		{
			Url:    parseURL("http://wetter-edka.de/cam2.jpg"),
			Source: wetterEdka,
		},
		{
			Url:    parseURL("http://wetter-edka.de/cam3l.jpg"),
			Source: wetterEdka,
		},
		{
			Url:    parseURL("http://wetter-edka.de/cam4l.jpg"),
			Source: wetterEdka,
		},
		{
			Url: parseURL("https://acdom-cdn.contentfux.de/live-webcam2.jpg"),
		},
		{
			Url: parseURL("http://www.zzz.at/webcams/pontstrasse/ponttor.cgi"),
		},
		{
			Url: parseURL("http://www.zzz.at/webcams/pontstrasse/pont.cgi"),
		},
		{
			Url: parseURL("http://www.salsatecas.de/z/webcams/eschweiler/cam24/esw1.cgi"),
		},
		{
			Url: parseURL("http://www.salsatecas.de/z/webcams/eschweiler/cam24/esw2.cgi"),
		},
		{
			Url: parseURL("http://www.salsatecas.net/z/webcams/eschweiler/cam26/WEBC-M12-ESW.jpg"),
		},
		{
			Url: parseURL("http://www.salsatecas.de/z/webcams/eschweiler/cam26/ESW-WEBC-ROEHE.jpg"),
		},
		{
			Url: parseURL("http://www.salsatecas.de/z/webcams/pont3/pont3_00001.jpg"),
		},
		{
			Url: parseURL("http://www.salsatecas.de/z/webcams/pont2/pont2.jpg"),
		},
		{
			Url: parseURL("http://www.caragh-lake.de/webcam/lafinestra/finex_00001.jpg"),
		},
		{
			Url:    parseURL("http://webcambild-rathaus.aachen.de/webcam_rathaus.jpg"),
			Source: parseURL("https://www.aachen.de/DE/stadt_buerger/aachen_profil/webcam/index.html"),
		},
		{
			Url:    parseURL("https://www.antenneac.de/webcam/studio01.jpg"),
			Source: parseURL("https://www.antenneac.de/ueber-uns/webcam.html"),
		},
		{
			Url:    parseURL("https://ak-aachen.de/webcam/uploads/Dome1.jpg"),
			Source: akAachen,
		},
		{
			Url:    parseURL("https://ak-aachen.de/webcam/uploads/Dome2.jpg"),
			Source: akAachen,
		},
		{
			Url:    parseURL("https://ak-aachen.de/webcam/uploads/Dome3.jpg"),
			Source: akAachen,
		},
		{
			Url:    parseURL("https://ak-aachen.de/webcam/uploads/Dome4.jpg"),
			Source: akAachen,
		},
	}
}

func GetWebcams() []Webcam {
	return Webcams
}
