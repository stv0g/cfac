package freifunk

// See https://data.aachen.freifunk.net/

// https://data.aachen.freifunk.net/nodes.json

// curl 'https://data.aachen.freifunk.net/meshviewer.json' \
//   -H 'sec-ch-ua: "Chromium";v="93", " Not;A Brand";v="99"' \
//   -H 'Referer: https://map.aachen.freifunk.net/' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'User-Agent: Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   --compressed

type NodeFlags struct {
	Online  bool `json:"online"`
	Gateway bool `json:"gateway"`
}

type NodeStatisticsProcesses struct {
	Total   int `json:"total"`
	Running int `json:"running"`
}

type NodeStatisticsTraffic struct {
	Tx struct {
		Bytes   int `json:"bytes"`
		Packets int `json:"packets"`
		Dropped int `json:"dropped"`
	} `json:"tx"`
	Rx struct {
		Bytes   int `json:"bytes"`
		Packets int `json:"packets"`
	} `json:"rx"`
	Forward struct {
		Bytes   int `json:"bytes"`
		Packets int `json:"packets"`
	} `json:"forward"`
	MgmtTx struct {
		Bytes   int `json:"bytes"`
		Packets int `json:"packets"`
	} `json:"mgmt_tx"`
	MgmtRx struct {
		Bytes   int `json:"bytes"`
		Packets int `json:"packets"`
	} `json:"mgmt_rx"`
}

type NodeStatisticsMeshVPN struct {
	Groups struct {
		Backbone struct {
			Peers struct {
				Aachen01 struct {
					Established float64 `json:"established"`
				} `json:"aachen01"`
				Aachen02 interface{} `json:"aachen02"`
				Aachen03 interface{} `json:"aachen03"`
				Aachen04 interface{} `json:"aachen04"`
				Aachen05 interface{} `json:"aachen05"`
				Aachen06 interface{} `json:"aachen06"`
			} `json:"peers"`
			Groups interface{} `json:"groups"`
		} `json:"backbone"`
	} `json:"groups"`
}

type NodeStatistics struct {
	NodeID      string                  `json:"node_id"`
	Clients     int                     `json:"clients"`
	RootfsUsage float64                 `json:"rootfs_usage"`
	Loadavg     float64                 `json:"loadavg"`
	MemoryUsage float64                 `json:"memory_usage"`
	Uptime      float64                 `json:"uptime"`
	Idletime    float64                 `json:"idletime"`
	Gateway     string                  `json:"gateway"`
	Gateway6    string                  `json:"gateway6"`
	Processes   NodeStatisticsProcesses `json:"processes"`
	MeshVpn     NodeStatisticsMeshVPN   `json:"mesh_vpn"`
	Traffic     NodeStatisticsTraffic   `json:"traffic"`
}

type NodeInfoNetwork struct {
	Mac       string   `json:"mac"`
	Addresses []string `json:"addresses"`
	Mesh      struct {
		Bat0 struct {
			Interfaces struct {
				Wireless []string `json:"wireless"`
				Other    []string `json:"other"`
				Tunnel   []string `json:"tunnel"`
			} `json:"interfaces"`
		} `json:"bat0"`
	} `json:"mesh"`
	MeshInterfaces interface{} `json:"mesh_interfaces"`
}

type NodeInfoSoftware struct {
	Autoupdater struct {
		Enabled bool   `json:"enabled"`
		Branch  string `json:"branch"`
	} `json:"autoupdater"`
	BatmanAdv struct {
		Version string `json:"version"`
		Compat  int    `json:"compat"`
	} `json:"batman-adv"`
	Babeld struct {
	} `json:"babeld"`
	Fastd struct {
		Enabled bool   `json:"enabled"`
		Version string `json:"version"`
	} `json:"fastd"`
	Firmware struct {
		Base    string `json:"base"`
		Release string `json:"release"`
	} `json:"firmware"`
	StatusPage struct {
		API int `json:"api"`
	} `json:"status-page"`
}

type NodeInfoHardware struct {
	Nproc int    `json:"nproc"`
	Model string `json:"model"`
}

type NodeInfoSystem struct {
	SiteCode string `json:"site_code"`
}

type NodeInfo struct {
	NodeID   string           `json:"node_id"`
	Network  NodeInfoNetwork  `json:"network"`
	System   NodeInfoSystem   `json:"system"`
	Hostname string           `json:"hostname"`
	Software NodeInfoSoftware `json:"software"`
	Hardware NodeInfoHardware `json:"hardware"`
	Vpn      bool             `json:"vpn"`
}

type Node struct {
	Firstseen  string         `json:"firstseen"`
	Lastseen   string         `json:"lastseen"`
	Flags      NodeFlags      `json:"flags"`
	Statistics NodeStatistics `json:"statistics"`
	Nodeinfo   NodeInfo       `json:"nodeinfo"`
}

type Nodes struct {
	Version   int    `json:"version"`
	Timestamp string `json:"timestamp"`
	Nodes     []Node `json:"nodes"`
}
