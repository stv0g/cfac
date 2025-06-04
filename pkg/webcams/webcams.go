// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

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

var Webcams []Webcam

func parseURL(u string) *url.URL {
	i, _ := url.Parse(u)
	return i
}

func init() {
	wetterEdka := parseURL("http://wetter-edka.de/")
	akAachen := parseURL("https://ak-aachen.de/webcams")
	verkehrnrw := parseURL("https://www.verkehr.nrw/")
	zzzat := parseURL("http://zzz.at/webcams/pontstrasse/")

	Webcams = []Webcam{
		// Flughafen Eschweiler
		{
			Title:      "Flughafen Eschweiler (EDKA)",
			SubTitle:   "LCS-1 --> W",
			StillImage: parseURL("http://wetter-edka.de/cam1.jpg"),
			Source:     wetterEdka,
			UpdateRate: time.Minute,
		},
		{
			Title:      "Flughafen Eschweiler (EDKA)",
			SubTitle:   "LCS-1 --> E",
			StillImage: parseURL("http://wetter-edka.de/cam2.jpg"),
			Source:     wetterEdka,
			UpdateRate: time.Minute,
		},
		{
			Title:      "Flughafen Eschweiler (EDKA)",
			StillImage: parseURL("http://wetter-edka.de/cam3l.jpg"),
			Source:     wetterEdka,
			UpdateRate: time.Minute,
		},
		{
			Title:      "Flughafen Eschweiler (EDKA)",
			StillImage: parseURL("http://wetter-edka.de/cam4l.jpg"),
			Source:     wetterEdka,
			UpdateRate: time.Minute,
		},

		// Aachener Dom
		{
			Title:      "Aachener Dom",
			SubTitle:   "Blick in den Falkenhorst Westturm",
			StillImage: parseURL("https://acdom-cdn.contentfux.de/live-webcam2.jpg"),
			Source:     parseURL("https://dombauhuette-aachen.de/webcam/webcam-innen/"),
			UpdateRate: 2 * time.Second,
		},
		{
			Title:      "Aachener Dom",
			SubTitle:   "Blick von der Laterne Oktogon auf Katschof und Quadrum",
			StillImage: parseURL("https://acdom-a-cdn.contentfux.de/live-webcam.jpg"),
			Source:     parseURL("https://dombauhuette-aachen.de/webcam/webcam-aussen/"),
			UpdateRate: 2 * time.Second,
		},

		// Pontstrasse
		{
			Title:      "Pontstrasse",
			SubTitle:   "Vlick in Richtung Ponttor",
			StillImage: parseURL("http://www.zzz.at/webcams/pontstrasse/ponttor.cgi"),
			Source:     zzzat,
		},
		{
			Title:      "Pontstrasse",
			StillImage: parseURL("http://www.zzz.at/webcams/pontstrasse/pont.cgi"),
			Source:     zzzat,
			UpdateRate: 3 * time.Minute,
		},
		{
			Title:      "Eschweiler",
			SubTitle:   "Blick von der Martin-Luther-Str. Richtung Norden über die Uferstrasse und die Indestrasse hinweg auf Kirche",
			StillImage: parseURL("http://www.salsatecas.de/z/webcams/eschweiler/cam24/esw1.cgi"),
			Source:     zzzat,
		},
		{
			Title:      "Eschweiler",
			SubTitle:   "Blick von der Martin-Luther-Str. Richtung Norden über die Uferstrasse und die Indestrasse hinweg auf Rathaus",
			StillImage: parseURL("http://www.salsatecas.de/z/webcams/eschweiler/cam24/esw2.cgi"),
			Source:     zzzat,
		},
		{
			Title:      "Eschweiler",
			SubTitle:   "Blickt in die Martin-Luther-Str. Richtung Marienstrasse und Krankenhaus",
			StillImage: parseURL("http://www.salsatecas.net/z/webcams/eschweiler/cam26/WEBC-M12-ESW.jpg"),
			LiveStream: parseURL("http://84.171.50.59:8026/record/current.jpg"),
			Source:     zzzat,
		},
		{
			Title:      "Eschweiler Röhe",
			SubTitle:   "Blick von Röhe auf das Zentrum von Eschweiler Richtung Osten/Südosten",
			StillImage: parseURL("http://www.salsatecas.de/z/webcams/eschweiler/cam26/ESW-WEBC-ROEHE.jpg"),
			Source:     zzzat,
		},
		{
			Title:      "Pontstrasse",
			SubTitle:   "La Finestra 3",
			StillImage: parseURL("http://www.salsatecas.de/z/webcams/pont3/pont3_00001.jpg"),
			Source:     zzzat,
		},
		{
			Title:      "Pontstrasse",
			SubTitle:   "La Finestra Erdgeschoss",
			StillImage: parseURL("http://www.salsatecas.de/z/webcams/pont2/pont2.jpg"),
			Source:     zzzat,
			UpdateRate: 20 * time.Minute,
		},
		{
			Title:      "Pontstrasse",
			SubTitle:   "La Finestra 2",
			StillImage: parseURL("http://www.caragh-lake.de/webcam/lafinestra/finex_00001.jpg"),
			Source:     zzzat,
		},

		// Filmpost
		{
			Title:      "Eschweiler",
			SubTitle:   "Filmpost Markt",
			LiveStream: parseURL("https://cam.filmpost.de/stream.m3u8"),
			Source:     parseURL("https://www.filmpost.de/"),
		},

		// Eifelwetter
		{
			Title:      "Wetterstation Monschau-Mützenich",
			SubTitle:   "Cam 1",
			StillImage: parseURL("https://eifelwetter.de/webcam/webcam1.jpg"),
			Source:     parseURL("https://www.eifelwetter.de/"),
		},
		{
			Title:      "Wetterstation Monschau-Mützenich",
			SubTitle:   "Cam 2",
			StillImage: parseURL("https://eifelwetter.de/webcam/webcam2.jpg"),
			Source:     parseURL("https://www.eifelwetter.de/"),
		},

		// AK Aachen
		{
			Title:      "A544",
			SubTitle:   "Blick nach Nordwest",
			StillImage: parseURL("https://ak-aachen.de/webcam/uploads/Dome1.jpg"),
			Source:     akAachen,
			UpdateRate: time.Minute,
		},
		{
			Title:      "A4",
			SubTitle:   "Blick nach Köln",
			StillImage: parseURL("https://ak-aachen.de/webcam/uploads/Dome2.jpg"),
			Source:     akAachen,
			UpdateRate: time.Minute,
		},
		{
			Title:      "Überfliegerbrücke A4 über A544",
			StillImage: parseURL("https://ak-aachen.de/webcam/uploads/Dome3.jpg"),
			Source:     akAachen,
			UpdateRate: time.Minute,
		},
		{
			Title:      "A544/A4",
			SubTitle:   "Blick nach AC-Europaplatz",
			StillImage: parseURL("https://ak-aachen.de/webcam/uploads/Dome4.jpg"),
			Source:     akAachen,
			UpdateRate: time.Minute,
		},

		// Verkehr NRW
		{
			Title:      "A4 | ID191 AK Aachen",
			SubTitle:   "Blickrichtung Köln",
			StillImage: parseURL("https://www.verkehr.nrw/webcams/10101803242723150838.jpg"),
			LiveStream: parseURL("https://www.blitzvideoserver.de/player_strassennrw.html?serverip=62.113.210.7&serverapp=strassennrw-rtplive&streamname=10101803242723150838"),
			Source:     verkehrnrw,
		},
		{
			Title:      "A4 | ID191 AK Aachen",
			SubTitle:   "Blickrichtung Köln",
			StillImage: parseURL("https://www.verkehr.nrw/webcams/10106382188859403675.jpg"),
			LiveStream: parseURL("https://www.blitzvideoserver.de/player_strassennrw.html?serverip=62.113.210.7&serverapp=strassennrw-rtplive&streamname=10106382188859403675"),
			Source:     verkehrnrw,
		},
		{
			Title:      "A544 | ID021 AK Aachen",
			SubTitle:   "Blickrichtung Aachen",
			StillImage: parseURL("https://www.verkehr.nrw/webcams/10107542537213142077.jpg"),
			LiveStream: parseURL("https://www.blitzvideoserver.de/player_strassennrw.html?serverip=62.113.210.7&serverapp=strassennrw-rtplive&streamname=10107542537213142077"),
			Source:     verkehrnrw,
		},

		{
			Title:      "A544 | ID021 AK Aachen",
			SubTitle:   "Blickrichtung Köln",
			StillImage: parseURL("https://www.verkehr.nrw/webcams/10107898794605103925.jpg"),
			LiveStream: parseURL("https://www.blitzvideoserver.de/player_strassennrw.html?serverip=62.113.210.7&serverapp=strassennrw-rtplive&streamname=10107898794605103925"),
			Source:     verkehrnrw,
		},

		// Other
		{
			Title:      "Rathaus Aachen",
			StillImage: parseURL("http://webcambild-rathaus.aachen.de/webcam_rathaus.jpg"),
			Source:     parseURL("https://www.aachen.de/DE/stadt_buerger/aachen_profil/webcam/index.html"),
		},
		{
			Title:      "Antenna AC Studio",
			StillImage: parseURL("https://www.antenneac.de/webcam/studio01.jpg"),
			Source:     parseURL("https://www.antenneac.de/ueber-uns/webcam.html"),
		},
	}
}

func GetWebcams() []Webcam {
	return Webcams
}
