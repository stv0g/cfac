package carolus

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (o Occupancy) Measure() []cfac.Measurement {
	base := func(name string) cfac.BaseMeasurement {
		return cfac.BaseMeasurement{
			Name: "carolus",
			Object: cfac.Object{
				Name:     "Carolus Therme " + name,
				Location: cfac.Coordinate{},
			},
		}
	}

	return []cfac.Measurement{
		cfac.OccupancyPercentMeasurement{
			BaseMeasurement: base("Parken"),

			Occupancy: o.Parking,
		},
		cfac.OccupancyPercentMeasurement{
			BaseMeasurement: base("Sauna"),

			Occupancy: o.Sauna,
		},
		cfac.OccupancyPercentMeasurement{
			BaseMeasurement: base("Thermalbad"),

			Occupancy: o.ThermalBath,
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchOccupancy(c, func(o Occupancy) {
		cb(o.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("carolus", NewMeasurable)
}
