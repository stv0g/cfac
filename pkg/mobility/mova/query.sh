#!/bin/bash

BBOX="50.88844,5.850671,50.666303,6.349731"
TYPES="BUS_STOP,CAR_SHARING_STATION,BIKE_SHARING_STATION,RAIL_STATION,POI"

HEADERS="-H 'Accept: application/json, text/plain, */*'
-H 'Referer: https://mova.aseag.de/'
-H 'X-MB-TENANT: ASEAG'
-H 'X-MB-AUTH-HEADER: qbsKtNX0hyjLxF2d3Z0qgA=='
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3857.0 Safari/537.36'
-H 'X-MB-FRONTEND-VERSION: 2.4.2'
-H 'Sec-Fetch-Mode: cors'
--compressed"

#curl "https://mova.aseag.de/mbroker/rest/areainformation?locationWindow=$BBOX&placeTypes=$TYPES" $HEADERS


curl "https://mova.aseag.de/mbroker/rest/freefloating?locationWindow=$BBOX&types=SCOOTER" $HEADERS
