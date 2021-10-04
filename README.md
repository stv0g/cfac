# cfac - Code for Aachen - Real-time APIs for Aachen

[![Go Reference](https://pkg.go.dev/badge/github.com/stv0g/cfac.svg)](https://pkg.go.dev/github.com/stv0g/cfac)
![](https://img.shields.io/snyk/vulnerabilities/github/stv0g/cfac)
[![](https://img.shields.io/github/checks-status/stv0g/cfac/master)](https://github.com/stv0g/cfac/actions)
[![](https://img.shields.io/librariesio/release/stv0g/cfac)](https://libraries.io/github/stv0g/cfac)
[![GitHub](https://img.shields.io/github/license/stv0g/cfac)](https://github.com/stv0g/cfac/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/stv0g/cfac)

_cfac_ is a Go package for accessing various real-time data sources and APIs in the [city of Aachen](https://aachen.de/).

## Disclaimer

_cfac_ is currently under active development and considered in an alpha stage.

_cfac_ accesses data from various sources by the means of web-scraping and undocumented APIs.

### Package contents

### Mobilität (`mobility`)

- [APAG](): Parkhäuser und aktuelle Auslastung
- [Velocity](): Stationen, Räder und Ladezustände
- [VOI](): Aktuelle Roller Standorte und Ladezustände
- [Cambio](): Stationen und aktuell verfügbare Mietwagen
- [Flinkster](): Stationen und aktuell verfügbare Mietwagen
- [Ford Carsharing](https://www.ford-carsharing.de/de/standorte?p=1)
- [MOVA](https://mova.aseag.de/#/home)
- [DB / Hafas]()
- [Flüge](https://de.flightaware.com/live/airport/EDKA): Aktuelle Flüge über Aachener Stadtgebiet
- [Baustelleninformationsssytem](): Aktuelle Baustellen im Aachener Stadtgebiet
    - http://aachen.de/DE/stadt_buerger/verkehr_strasse/strassenplanung_bau/bsis/index.html
    - https://bsis.aachen.de/
- [E-Ladesäulen]()

## Besucherzahlen & aktuelle Auslastung (`occupancy`)

- [Carolus Therme](https://www.carolus-thermen.de/en/thermalbath/#occupation)
- [Bürgerservice Wartezeiten](https://serviceportal.aachen.de/wartezeiten)
- [Chaos Computer Club Aachen](https://wiki.aachen.ccc.de/doku.php?id=projekte:clubstatus)
- Fitnesscenter:
  - [FitX](https://www.fitx.de/fitnessstudios/aachen-europaplatz)
  - [McFit](https://www.mcfit.com/de/fitnessstudios/studiosuche/studiodetails/studio/aachen/)
  - [RWTH Hochschulsport](https://buchung.hsz.rwth-aachen.de/angebote/aktueller_zeitraum/_Auslastung.html)
  - [MedAix](https://www.medaix.de/standorte/aachen-elisengalerie)

### Wetter (`weather`)

- [AWEKAS](https://www.awekas.at/)
- [meteo Aachen](https://meteoaachen.de/)
- [Deutscher Wetterdienst](https://www.dwd.de/DE/wetter/wetterundklima_vorort/nordrhein-westfalen/aachen/_node.html)
- [Eifelwetter](https://www.eifelwetter.de/): aktuelle Wetterdaten
- [LANUV](https://www.lanuv.nrw.de/umwelt/luft/immissionen/aktuelle-luftqualitaet): aktuelle Luftqualität
- [ODLinfo](https://odlinfo.bfs.de/DE/aktuelles/messstelle/053130003.html): aktuelle Radioaktivitäts / ODL Messungen

### Webcams (`webcams`)

### Gesundheit (`health`)

- [UK Aachen](https://www.ukaachen.de/kliniken-institute/transfusionsmedizin-blutspendedienst/blutspendedienst/blutspendepegel/spendepegel/2021-08): aktueller Blutspendepegel 
- [Covid 19](): aktuelle Fallzahlen in der Städteregion

### Sonstiges

- [Abfallapp/navi](https://abfallnavi.de/aachen/)
- [Freifunk Aachen](https://map.aachen.freifunk.net/): Standorte Accesspoints, aktuelle Auslastung, [Open-data](https://data.aachen.freifunk.net)
- [RWTH Aachen University](): aktuelle Auslastung Lernräume

## Roadmap: future ideas

### Opening hours

- Lieferando: 
- Lieferheld: https://github.com/kenodressel/lieferheld-api/blob/master/info-api.js

## Links

- https://bund.dev/
- https://github.com/bundesAPI
- http://opendata.aachen.de/
- https://geoportal.aachen.de/
- https://serviceportal.aachen.de/abfallnavi
- http://ratsinfo.aachen.de/bi/allris.net.asp
