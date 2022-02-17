# Verkehrs Dashboard

- https://verkehr.aachen.de
- https://verkehr.aachen.de/config/dashboard

## SONAH

- https://aachen.digital/news/best-practice-sonah-aachen/
- https://sonah.tech/index_main.html
- https://verkehr.aachen.de/api/sonah/api/v2/locations?include=Name,TotalParking,FreeParking,Type,Positions,ParentLocations

## HAFAS

- https://verkehr.aachen.de/api/hafas/departureBoard?id=A=1@O=Aachen,%20Elisenbrunnen@X=6087513@Y=50774158@u=357@U=80@L=1029@&rtMode=FULL&duration=20&date=2022-02-17&time=19:08
- https://verkehr.aachen.de/api/hafas/location.nearbystops?originCoordLong=6.083937&originCoordLat=50.776444&r=11300&type=S&maxNo=1000

## Mobility Broker / MOVAS

- https://www.bettermobility.de/produkte/#mbb
- https://verkehr.aachen.de/api/mbroker/areainformation?locationWindow=50.8219872,6.0071182,50.7203733,6.1976624&placeOrOtherMobilityTypes=MOBILITAETSSTATION,CAR_SHARING_STATION,BIKE_SHARING_STATION,CAR_SHARING,BIKE_SHARING
- https://verkehr.aachen.de/api/mbroker/areainformation/29872

## SensorThings

- Parkobjekt APAG: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27Parkobjekt%20APAG%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,Datastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters),Datastreams%2FObservedProperty(%24select%3D%40iot.id%2Cname)&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name
- Wetter: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=name%20eq%20%2710501%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,Datastreams,Datastreams/ObservedProperty,Datastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters)
- RydeUp: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27Rydeup%27%20and%20properties/archive%20eq%20%27false%27&&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name
- ChargingStations: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27E-Ladestation%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations(%24select%3D%40iot.id%2Clocation),Datastreams(%24select%3D%40iot.id%2Cname%2Cproperties),Datastreams%2FObservedProperty(%24select%3D%40iot.id%2Cname),Datastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters)&$top=300&$select=@iot.id,description,name,properties&$orderBy=description
- Infrarot Fahrradzähler: https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27Verkehrszaehlstelle%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,Datastreams(%24filter%3Dproperties%2FKlasse%20eq%20%27Bike%27%20and%20properties%2FAggregation%20eq%20%27d%27),Datastreams%2FObservedProperty,Datastreams%2FObservations(%24top%3D7%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime)&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name

## Verkehr / NRW

- https://www.verkehr.nrw/web/vipnrw/karte?p_p_id=de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_cacheability=cacheLevelPage&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_car=true&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_city=aachen&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_center=50.77499%2C6.08360&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_highlightRoute=false&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_zoom=14&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_layer=Verkehrslage%2CVerkehrsmeldungen&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_publicTransport=true&_de_strassennrw_vipnrw_portlet_news_portlet_NewsPortlet_bike=true

## MasterPortal

- https://bitbucket.org/geowerkstatt-hamburg/masterportalapi/src/master/

- https://verkehr-mp.aachen.de/FROST-Server/v1.1/Things
- https://verkehr-mp.aachen.de/resources/services-internet.json
- https://verkehr-mp.aachen.de/config.json
