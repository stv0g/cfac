// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package mova

const (
	UrlApi             = "https://mova.aseag.de/mbroker/rest"
	UrlAreaInformation = UrlApi + "/areainformation?locationWindow=50.772693748685995,6.080337166786195,50.77139442220511,6.089000701904297&placeOrOtherMobilityTypes=BUS_STOP,MOBILITAETSSTATION,CAR_SHARING_STATION,BIKE_SHARING_STATION,RAIL_STATION,POI,ADDRESS,BUS,CAR_SHARING,BIKE_SHARING,RAILWAY"
	UrlFreeFloating    = UrlApi + "/freefloating?locationWindow=50.772693748685995,6.080337166786195,50.77139442220511,6.089000701904297&types=SCOOTER,BIKE_SHARING,CAR_SHARING"
)
