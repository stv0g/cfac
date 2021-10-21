package warnung

const (
	UrlApi          = "https://warnung.bund.de/api31"
	UrlApiDashboard = UrlApi + "/dashboard/{ags}.json"
	UrlApiKatwarn   = UrlApi + "/katwarn/mapData.json"
	UrlApiBiwapp    = UrlApi + "/biwapp/mapData.json"
	UrlApiMowas     = UrlApi + "/mowas/mapData.json"
	UrlApiAppCovid  = UrlApi + "/appdata/covid/covidrules/DE/{ags}.json"

	AGSAachen = "053340000000"
)
