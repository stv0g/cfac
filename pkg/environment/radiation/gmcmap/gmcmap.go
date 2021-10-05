package gmcmap

// curl -o complete-gmcmap-mapdevices.json https://www.gmcmap.com/AJAX_load_time.asp\?OffSet\=0\&Limit\=2147483647\&dataRange\=100000\&timeZone\=2

const (
	Url            = "http://www.gmcmap.com"
	UrlPlainData   = Url + "/gmc-plain-data.asp"    // + ?Param_ID=YOUR_GEIGER_COUNTER_ID&timezone=TIME_ZONE
	UrlHistoryData = Url + "/historyData-plain.asp" // + ?Param_ID=64609518812&timezone=-1
)

var (
	SensorIDs = []int{
		53122006742,
	}
)
