# cfac - Code for Aachen - Real-time APIs for Aachen

[![Go Reference](https://pkg.go.dev/badge/github.com/stv0g/cfac.svg)](https://pkg.go.dev/github.com/stv0g/cfac)
![Snyk](https://img.shields.io/snyk/vulnerabilities/github/stv0g/cfac)
[![Build](https://img.shields.io/github/checks-status/stv0g/cfac/master)](https://github.com/stv0g/cfac/actions)
[![libraries.io](https://img.shields.io/librariesio/release/stv0g/cfac)](https://libraries.io/github/stv0g/cfac)
[![GitHub](https://img.shields.io/github/license/stv0g/cfac)](https://github.com/stv0g/cfac/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/stv0g/cfac)

_cfac_ is a Go package for accessing various real-time data sources and APIs in the [city of Aachen](https://aachen.de/).

## Disclaimer

_cfac_ is currently under active development and considered in an alpha stage.

_cfac_ accesses data from various sources by the means of web-scraping and undocumented APIs.

## Package contents

### Aachener Verkehrs Dashboard

- [APIs](VerkehrsDashboard.md)

### Mobilität (`mobility`)

- ✅ [APAG](https://www.apag.de/): Parkhäuser und aktuelle Auslastung
- [Velocity](https://www.velocity-aachen.de/): Stationen, Räder und Ladezustände
- [VOI](https://www.voiscooters.com/de/): Aktuelle Roller Standorte und Ladezustände
- [Cambio](https://www.cambio-carsharing.de/?cms_knschluessel=HOME&cms_Feurocode=AAC): Stationen und aktuell verfügbare Mietwagen
- [Flinkster](https://www.flinkster.de/): Stationen und aktuell verfügbare Mietwagen
- [Ford Carsharing](https://www.ford-carsharing.de/de/standorte?p=1)
- [MOVA](https://mova.aseag.de/)
- [Deutsche Bahn (HAFAS)](https://github.com/public-transport/hafas-client/blob/master/readme.md#background)
- [Flüge](https://de.flightaware.com/live/airport/EDKA): Aktuelle Flüge über Aachener Stadtgebiet
- [Baustelleninformationsssytem](https://bsis.aachen.de/): Aktuelle Baustellen im Aachener Stadtgebiet
  - [Beschreibung](http://aachen.de/DE/stadt_buerger/verkehr_strasse/strassenplanung_bau/bsis/index.html)

#### Elektro Ladeinfrastruktur (`ev_charger`)

- [Open Charge Map](https://openchargemap.org/site)
- [plugsurfing](https://www.plugsurfing.com/map?location=Aachen,%20Germany&lang=en)

### Besucherzahlen & aktuelle Auslastung (`occupancy`)

- ✅ [Carolus Therme](https://www.carolus-thermen.de/en/thermalbath/#occupation)
- ✅ [Spielbank Aachen](https://www.spielbank-aachen.de)
- ✅ [Chaos Computer Club Aachen](https://wiki.aachen.ccc.de/doku.php?id=projekte:clubstatus)

#### Fitnessstudios (`gyms`)

- ✅ [FitX](https://www.fitx.de/fitnessstudios/aachen-europaplatz)
- ✅ [McFit](https://www.mcfit.com/de/fitnessstudios/studiosuche/studiodetails/studio/aachen/)
- ✅ [RWTH Hochschulsport](https://buchung.hsz.rwth-aachen.de/angebote/aktueller_zeitraum/_Auslastung.html)
- ✅ [MedAix](https://www.medaix.de/standorte/aachen-elisengalerie)
- ✅ [Word of Fitness](http://besucher.wof-fitness.de/)

### Umwelt (`environment`)

- [sensor.community](https://sensor.community)

#### Wetter (`weather`)

- [AWEKAS](https://www.awekas.at/)
- [meteo Aachen](https://meteoaachen.de/)
- [Deutscher Wetterdienst](https://www.dwd.de/DE/wetter/wetterundklima_vorort/nordrhein-westfalen/aachen/_node.html)
- [Eifelwetter](https://www.eifelwetter.de/): aktuelle Wetterdaten
- [Buienalarm](https://buienalarm.nl)

#### Luftqualität (`air_quality`)

- [LANUV](https://www.lanuv.nrw.de/umwelt/luft/immissionen/aktuelle-luftqualitaet): aktuelle Luftqualität

#### Strahlung (`radiation`)

- [ODLinfo](https://odlinfo.bfs.de/DE/aktuelles/messstelle/053130003.html): aktuelle Radioaktivitäts / ODL Messungen
- [Gammasense](https://gammasense.org/map/)
- [Geige Counter World Map (gmcmap.com)](https://www.gmcmap.com/)
- [radmon.org](https://radmon.org/index.php)
- [Tihange-Doel Radiation Monitoring (TDRM)](https://tdrm.fiff.de/)

### Webcams (`webcams`)

### Gesundheit (`health`)

- ✅ [Blutspendepegel](https://www.ukaachen.de/kliniken-institute/transfusionsmedizin-blutspendedienst/blutspendedienst/blutspendepegel/spendepegel/2021-08)
- [Covid 19](https://offenedaten.aachen.de/dataset/aktuelle-lage-zum-corona-virus): aktuelle Fallzahlen in der Städteregion

### Sonstiges

- [Abfallapp/navi](https://abfallnavi.de/aachen/)
- [Freifunk Aachen](https://map.aachen.freifunk.net/): Standorte Accesspoints, aktuelle Auslastung, [Open-data](https://data.aachen.freifunk.net)

- [Amateuerfunk-Gruppe der RWTH Aachen](https://www.afu.rwth-aachen.de/)
  - APRS Receiver

- [RTL-SDR](https://www.rtl-sdr.com/about-rtl-sdr/)
  - Tracking aircraft positions like a radar with [ADSB decoding](https://www.rtl-sdr.com/adsb-aircraft-radar-with-rtl-sdr/)
  - [Receiving NOAA weather satellite images.](https://www.rtl-sdr.com/rtl-sdr-tutorial-receiving-noaa-weather-satellite-images/)

### Out-of-service

**TODO:** Contact data providers via mail..

- [Wetterstation am Physikzentrum der RWTH Aachen](https://wetterstation.physik.rwth-aachen.de/)
- [Wetterstation Geo RWTH](https://www.klimageo.rwth-aachen.de/cms/Klimageo/Das-Lehr-und-Forschungsgebiet/Ausstattung/~pcdx/Wetterstationen/)
- [RWTH Aachen Lernraumampel](https://blog.rwth-aachen.de/lehre/2017/07/28/die-lernraumampel-ist-da/)
- [Bürgerservice Wartezeiten](https://serviceportal.aachen.de/wartezeiten)

## Roadmap: future ideas

- [Ratsinformations System Aachen](https://ratsinfo.aachen.de/bi/oparl/1.0/bodies.asp?id=1)
  - Statistiken..

### Opening hours

- Lieferando:
- Lieferheld
  - [API example](https://github.com/kenodressel/lieferheld-api/blob/master/info-api.js)

## Links

- [bund.dev](https://bund.dev)
- [bundesAPI](https://github.com/bundesAPI)
- Aachen
  - [Open Data](https://offenedaten.aachen.de)
  - [GeoPortal](https://geoportal.aachen.de)
  - [Ratsinfo](http://ratsinfo.aachen.de/bi/allris.net.asp)
