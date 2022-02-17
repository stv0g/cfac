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

- https://verkehr.aachen.de/config/dashboard

#### SONAH

- https://aachen.digital/news/best-practice-sonah-aachen/
- https://sonah.tech/index_main.html
- https://verkehr.aachen.de/api/sonah/api/v2/locations?include=Name,TotalParking,FreeParking,Type,Positions,ParentLocations

#### HAFAS

- https://verkehr.aachen.de/api/hafas/departureBoard?id=A=1@O=Aachen,%20Elisenbrunnen@X=6087513@Y=50774158@u=357@U=80@L=1029@&rtMode=FULL&duration=20&date=2022-02-17&time=19:08
- https://verkehr.aachen.de/api/hafas/location.nearbystops?originCoordLong=6.083937&originCoordLat=50.776444&r=11300&type=S&maxNo=1000

#### Mobility Broker / MOVAS

- https://www.bettermobility.de/produkte/#mbb
- https://verkehr.aachen.de/api/mbroker/areainformation?locationWindow=50.8219872,6.0071182,50.7203733,6.1976624&placeOrOtherMobilityTypes=MOBILITAETSSTATION,CAR_SHARING_STATION,BIKE_SHARING_STATION,CAR_SHARING,BIKE_SHARING
- https://verkehr.aachen.de/api/mbroker/areainformation/29872

#### SensorThings

- Parkobjekt APAG: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27Parkobjekt%20APAG%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,Datastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters),Datastreams%2FObservedProperty(%24select%3D%40iot.id%2Cname)&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name
- Wetter: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=name%20eq%20%2710501%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,Datastreams,Datastreams/ObservedProperty,Datastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters)
- RydeUp: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27Rydeup%27%20and%20properties/archive%20eq%20%27false%27&&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name
- ChargingStations: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27E-Ladestation%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations(%24select%3D%40iot.id%2Clocation),Datastreams(%24select%3D%40iot.id%2Cname%2Cproperties),Datastreams%2FObservedProperty(%24select%3D%40iot.id%2Cname),Datastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters)&$top=300&$select=@iot.id,description,name,properties&$orderBy=description
- Infrarot Fahrradzähler: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27Verkehrszaehlstelle%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,Datastreams(%24filter%3Dproperties%2FKlasse%20eq%20%27Bike%27%20and%20properties%2FAggregation%20eq%20%27d%27),Datastreams%2FObservedProperty,Datastreams%2FObservations(%24top%3D7%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime)&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name

#### Verkehr / NRW

- https://www.verkehr.nrw/web/vipnrw/karte?p_p_id=de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_cacheability=cacheLevelPage&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_car=true&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_city=aachen&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_center=50.77499%2C6.08360&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_highlightRoute=false&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_zoom=14&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_layer=Verkehrslage%2CVerkehrsmeldungen&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_publicTransport=true&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_bike=true

#### MasterPortal

- https://bitbucket.org/geowerkstatt-hamburg/masterportalapi/src/master/

- https://verkehr-mp.aachen.de/FROST-Server/v1.1/Things
- https://verkehr-mp.aachen.de/resources/services-internet.json
- https://verkehr-mp.aachen.de/config.json

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
