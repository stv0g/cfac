package freifunk

import (
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct{}

func (n *NodeNodelist) Measure() cfac.Measurement {
	var coord *cfac.Coordinate = nil
	if n.Position.Lat != 0 && n.Position.Long != 0 {
		coord = &cfac.Coordinate{
			Latitude:  n.Position.Lat,
			Longitude: n.Position.Long,
		}
	}

	return &cfac.CounterMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "counter",
			Metric: "clients",
			Source: "freifunk",
			Object: &cfac.Object{
				Name:     n.Name,
				Location: coord,
			},
			Time: uint64(n.Status.LastContact.UnixMilli()),
		},

		Count: uint64(n.Status.Clients),
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchNodeList(c, func(nl ResponseNodeList) {
		total := len(nl.Nodes)
		online := 0
		clients := 0

		for _, node := range nl.Nodes {
			cb(node.Measure())

			clients += node.Status.Clients

			if node.Status.Online {
				online++
			}
		}

		cb(&cfac.CounterMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "counter",
				Metric: "nodes_online",
				Source: "freifunk",
				Time:   uint64(nl.UpdatedAt.UnixMilli()),
			},
			Count: uint64(online),
		})

		cb(&cfac.CounterMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "counter",
				Metric: "nodes_offline",
				Source: "freifunk",
				Time:   uint64(nl.UpdatedAt.UnixMilli()),
			},
			Count: uint64(total - online),
		})

		cb(&cfac.CounterMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "counter",
				Metric: "clients_total",
				Source: "freifunk",
				Time:   uint64(nl.UpdatedAt.UnixMilli()),
			},
			Count: uint64(clients),
		})
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("freifunk", NewMeasurable)
}
