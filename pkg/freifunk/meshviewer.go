package freifunk

import "time"

// See https://data.aachen.freifunk.net/

const (
	Url = "https://data.aachen.freifunk.net/"

	UrlNodelist          = Url + "/nodelist.json"
	UrlNodes             = Url + "/nodes.json"
	UrlGraph             = Url + "/graph.json"
	UrlMeshviewer        = Url + "/meshviewer.json"
	UrlDistricStatistics = Url + "/ffac-district-statistics.json"
)

type ResponseNodes struct {
	Version   int    `json:"version"`
	Timestamp string `json:"timestamp"`
	Nodes     []Node `json:"nodes"`
}

type ResponseNodeList struct {
	Version   string         `json:"version"`
	UpdatedAt string         `json:"updated_at"`
	Nodes     []NodeNodelist `json:"nodes"`
}

type ResponseMeshviewer struct {
	Timestamp string           `json:"timestamp"`
	Nodes     []NodeMeshviewer `json:"nodes"`
	Links     []Link           `json:"links"`
}

type ResponseGraph struct {
	Version int `json:"version"`
	Batadv  struct {
		Directed bool        `json:"directed"`
		Graph    interface{} `json:"graph"`
		Nodes    []struct {
			ID     string `json:"id"`
			NodeID string `json:"node_id"`
		} `json:"nodes"`
		Links []struct {
			Source   int  `json:"source"`
			Target   int  `json:"target"`
			Vpn      bool `json:"vpn"`
			Tq       int  `json:"tq"`
			Bidirect bool `json:"bidirect"`
		} `json:"links"`
	} `json:"batadv"`
}

type ResponseDistricStatistics struct {
	Counts struct {
		KnownWithinBoundary  int `json:"known_within_boundary"`
		OnlineWithinBoundary int `json:"online_within_boundary"`
		WithGeo              int `json:"with_geo"`
		WithoutGeo           int `json:"without_geo"`
	} `json:"counts"`
	Districts struct {
		Known  map[string]int `json:"known"`
		Online map[string]int `json:"online"`
	} `json:"districts"`
	Timestamp time.Time `json:"timestamp"`
}

type NodeNodelist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status struct {
		Online      bool   `json:"online"`
		Lastcontact string `json:"lastcontact"`
		Clients     int    `json:"clients"`
	} `json:"status"`
	Position struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"position,omitempty"`
}

type Link struct {
	Type       string  `json:"type"`
	Source     string  `json:"source"`
	Target     string  `json:"target"`
	SourceTq   float64 `json:"source_tq"`
	TargetTq   float64 `json:"target_tq"`
	SourceAddr string  `json:"source_addr"`
	TargetAddr string  `json:"target_addr"`
}

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

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type NodeMeshviewer struct {
	Firstseen      string   `json:"firstseen"`
	Lastseen       string   `json:"lastseen"`
	IsOnline       bool     `json:"is_online"`
	IsGateway      bool     `json:"is_gateway"`
	Clients        int      `json:"clients"`
	ClientsWifi24  int      `json:"clients_wifi24"`
	ClientsWifi5   int      `json:"clients_wifi5"`
	ClientsOther   int      `json:"clients_other"`
	RootfsUsage    float64  `json:"rootfs_usage"`
	Loadavg        int      `json:"loadavg"`
	MemoryUsage    float64  `json:"memory_usage"`
	Uptime         string   `json:"uptime"`
	GatewayNexthop string   `json:"gateway_nexthop,omitempty"`
	Gateway        string   `json:"gateway,omitempty"`
	Gateway6       string   `json:"gateway6,omitempty"`
	NodeID         string   `json:"node_id"`
	Mac            string   `json:"mac"`
	Addresses      []string `json:"addresses"`
	Domain         string   `json:"domain"`
	Hostname       string   `json:"hostname"`
	Firmware       struct {
		Base    string `json:"base"`
		Release string `json:"release"`
	} `json:"firmware"`
	Autoupdater struct {
		Enabled bool   `json:"enabled"`
		Branch  string `json:"branch"`
	} `json:"autoupdater"`
	Nproc    int      `json:"nproc"`
	Model    string   `json:"model,omitempty"`
	Location Location `json:"location,omitempty"`
}
