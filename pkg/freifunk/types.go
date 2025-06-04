// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package freifunk

import (
	"strings"
	"time"
)

type ResponseNodes struct {
	Version   int       `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Nodes     []Node    `json:"nodes"`
}

type ResponseNodeList struct {
	Version   string         `json:"version"`
	UpdatedAt CustomTime     `json:"updated_at"`
	Nodes     []NodeNodelist `json:"nodes"`
}

type ResponseMeshviewer struct {
	Timestamp time.Time        `json:"timestamp"`
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
		Online      bool       `json:"online"`
		LastContact CustomTime `json:"lastcontact"`
		Clients     int        `json:"clients"`
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
	Tx      TrafficStatistics `json:"tx"`
	Rx      TrafficStatistics `json:"rx"`
	Forward TrafficStatistics `json:"forward"`
	MgmtTx  TrafficStatistics `json:"mgmt_tx"`
	MgmtRx  TrafficStatistics `json:"mgmt_rx"`
}

type TrafficStatistics struct {
	Bytes   int `json:"bytes"`
	Packets int `json:"packets"`
	Dropped int `json:"dropped"`
}

type NodeStatisticsMeshVPN struct {
	Groups struct {
		Backbone struct {
			Peers  map[string]PeerStatistics `json:"peers"`
			Groups map[string]PeerStatistics `json:"groups"`
		} `json:"backbone"`
	} `json:"groups"`
}

type PeerStatistics struct {
	Established float64 `json:"established"`
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
	AutoUpdater struct {
		Enabled bool   `json:"enabled"`
		Branch  string `json:"branch"`
	} `json:"autoupdater"`
	BatmanAdv struct {
		Version string `json:"version"`
		Compat  int    `json:"compat"`
	} `json:"batman-adv"`
	Babeld struct{} `json:"babeld"`
	Fastd  struct {
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
	Firstseen      string      `json:"firstseen"`
	Lastseen       string      `json:"lastseen"`
	IsOnline       bool        `json:"is_online"`
	IsGateway      bool        `json:"is_gateway"`
	Clients        int         `json:"clients"`
	ClientsWifi24  int         `json:"clients_wifi24"`
	ClientsWifi5   int         `json:"clients_wifi5"`
	ClientsOther   int         `json:"clients_other"`
	RootfsUsage    float64     `json:"rootfs_usage"`
	Loadavg        int         `json:"loadavg"`
	MemoryUsage    float64     `json:"memory_usage"`
	Uptime         string      `json:"uptime"`
	GatewayNexthop string      `json:"gateway_nexthop,omitempty"`
	Gateway        string      `json:"gateway,omitempty"`
	Gateway6       string      `json:"gateway6,omitempty"`
	NodeID         string      `json:"node_id"`
	Mac            string      `json:"mac"`
	Addresses      []string    `json:"addresses"`
	Domain         string      `json:"domain"`
	Hostname       string      `json:"hostname"`
	Firmware       Firmware    `json:"firmware"`
	AutoUpdater    AutoUpdater `json:"autoupdater"`
	Nproc          int         `json:"nproc"`
	Model          string      `json:"model,omitempty"`
	Location       Location    `json:"location,omitempty"`
}

type Firmware struct {
	Base    string `json:"base"`
	Release string `json:"release"`
}

type AutoUpdater struct {
	Enabled bool   `json:"enabled"`
	Branch  string `json:"branch"`
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.Parse("2006-01-02T15:04:05-0700", s)
	return err
}
