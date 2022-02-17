package sensorthings

// OGC SensorThings API Part 1: Sensing Version 1.1
// See: https://docs.ogc.org/is/18-088/18-088.html

import (
	"time"
)

type Repsonse struct {
	IotNextLink string  `json:"@iot.nextLink"`
	Value       []Thing `json:"value"`
}

type Any interface{}
type ValueCode string
type IotNavigationLink string

// A JSON Object containing user-annotated properties as key-value pairs.
type Properties map[string]Any

// In SensorThings control information is represented as annotations whose names start with iot followed by a dot ( . ).
// Annotations are name/value pairs that have a dot ( . ) as part of the name.
// See: https://docs.ogc.org/is/18-088/18-088.html#common-control-information
type Iot struct {
	// ID is the system-generated identifier of an entity.
	// ID is unique among the entities of the same entity type in a SensorThings service.
	ID int `json:"@iot.id"`

	// SelfLink is the absolute URL of an entity that is unique among all other entities.
	SelfLink IotNavigationLink `json:"@iot.selfLink"`
}

// The OGC SensorThings API follows the ITU-T definition, i.e., with regard to the Internet of Things, a thing is an object of the physical world (physical things) or the information world (virtual things) that is capable of being identified and integrated into communication networks [ITU-T Y.2060].
// See: https://docs.ogc.org/is/18-088/18-088.html#thing
type Thing struct {
	Iot

	// A property provides a label for Thing entity, commonly a descriptive name.
	Name string `json:"name"`

	// This is a short description of the corresponding Thing entity.
	Description string `json:"description"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	Locations          []Location            `json:"Locations"`
	Datastreams        []Datastream          `json:"Datastreams"`
	HistoricalLocation []HistoricalLocations `json:"HistoricalLocations"`

	LocationsLink           IotNavigationLink `json:"Locations@iot.navigationLink"`
	TaskingCapabilitiesLink IotNavigationLink `json:"TaskingCapabilities@iot.navigationLink"`
	HistoricalLocationsLink IotNavigationLink `json:"HistoricalLocations@iot.navigationLink"`
	MultiDatastreamsLink    IotNavigationLink `json:"MultiDatastreams@iot.navigationLink"`
	DatastreamsLink         IotNavigationLink `json:"Datastreams@iot.navigationLink"`
}

// The Location entity locates the Thing or the Things it associated with. A Thing’s Location entity is defined as the last known location of the Thing.
// See: https://docs.ogc.org/is/18-088/18-088.html#location
type Location struct {
	Iot

	// A property provides a label for Location entity, commonly a descriptive name.
	Name string `json:"name"`

	// The description about the Location.
	Description string `json:"description"`

	// The encoding type of the Location property. Its value is one of the ValueCode enumeration.
	// See: https://docs.ogc.org/is/18-088/18-088.html#tab-encodingtype-codes
	EncodingType string `json:"encodingType"`

	// The location type is defined by encodingType.
	Location GeoJSONLocation `json:"location"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	HistoricalLocations []HistoricalLocations `json:"HistoricalLocations"`
	Things              []Thing               `json:"Things"`

	HistoricalLocationsLink IotNavigationLink `json:"HistoricalLocations@iot.navigationLink"`
	ThingsLink              IotNavigationLink `json:"Things@iot.navigationLink"`
}

// A Thing’s HistoricalLocation entity set provides the times of the current (i.e., last known) and previous locations of the Thing.
// See: https://docs.ogc.org/is/18-088/18-088.html#historicallocation
type HistoricalLocations struct {
	Iot

	Time time.Time `json:"time"`

	Locations []Location `json:"Locations"`
	Thing     *Thing     `json:"Thing"`

	LocationsLink IotNavigationLink `json:"Locations@iot.navigationLink"`
	ThingLink     IotNavigationLink `json:"Thing@iot.navigationLink"`
}

// A Datastream groups a collection of Observations measuring the same ObservedProperty and produced by the same Sensor.
// See: https://docs.ogc.org/is/18-088/18-088.html#datastream
type Datastream struct {
	Iot

	// A property provides a label for Datastream entity, commonly a descriptive name.
	Name string `json:"name"`

	// The description of the Datastream entity.
	Description string `json:"description"`

	UnitOfMeasurement Unit `json:"unitOfMeasurement"`

	// The type of Observation (with unique result type), which is used by the service to encode observations.
	ObservationType ValueCode `json:"observationType"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	// The spatial bounding box of the spatial extent of all FeaturesOfInterest that belong to the Observations associated with this Datastream.
	ObservedArea GeoJSONPolygon `json:"observedArea"`

	// The temporal interval of the phenomenon times of all observations belonging to this Datastream.
	PhenomenonTime time.Time `json:"phenomenonTime"`

	// The temporal interval of the result times of all observations belonging to this Datastream.
	ResultTime time.Time `json:"resultTime"`

	// The Observations of a Datastream SHALL observe the same ObservedProperty. The Observations of different Datastreams MAY observe the same ObservedProperty.
	ObservedProperty ObservedProperty `json:"ObservedProperty"`

	Observations []Observation `json:"Observations"`

	Thing  *Thing  `json:"Thing"`
	Sensor *Sensor `json:"Sensor"`

	ThingLink        IotNavigationLink `json:"Thing@iot.navigationLink"`
	ObservationsLink IotNavigationLink `json:"Obversvation@iot.navigationLink"`
}

// A MultiDatastream groups a collection of Observations and the Observations in a MultiDatastream have a complex result type.
// See: https://docs.ogc.org/is/18-088/18-088.html#multidatastream-extension
type MultiDatastream struct {
	Iot

	// A property provides a label for Datastream entity, commonly a descriptive name.
	Name string `json:"name"`

	// The description of the Datastream entity.
	Description string `json:"description"`

	UnitOfMeasurements []Unit `json:"unitOfMeasurements"`

	// The type of Observation (with unique result type), which is used by the service to encode observations.
	ObservationType ValueCode `json:"observationType"`

	// This property defines the observationType of each element of the result of a complex Observation.
	MultiObservationDataTypes []ValueCode `json:"multiObservationDataTypes"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	// The spatial bounding box of the spatial extent of all FeaturesOfInterest that belong to the Observations associated with this Datastream.
	ObservedArea GeoJSONPolygon `json:"observedArea"`

	// The temporal interval of the phenomenon times of all observations belonging to this Datastream.
	PhenomenonTime time.Time `json:"phenomenonTime"`

	// The temporal interval of the result times of all observations belonging to this Datastream.
	ResultTime time.Time `json:"resultTime"`
}

//
type Unit struct {
	// Full name of unit of measurement
	Name string `json:"name"`

	// Textual form of the unit symbol
	Symbol interface{} `json:"symbol"`

	// URI to definition of unit
	Definition interface{} `json:"definition"`
}

// A Sensor is an instrument that observes a property or phenomenon with the goal of producing an estimate of the value of the property.
// See: https://docs.ogc.org/is/18-088/18-088.html#sensor
type Sensor struct {
	Iot

	// A property provides a label for ObservedProperty entity, commonly a descriptive name.
	Name string `json:"name"`

	// The description of the Sensor entity.
	Definition string `json:"definition"`

	// The encoding type of the metadata property. Its value is one of the ValueCode enumeration (see Table 15 for the available ValueCode).
	EncodingType ValueCode `json:"valueCode"`

	// The detailed description of the Sensor or system. The metadata type is defined by encodingType.
	Metadata Any `json:"metadata"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	Datastreams        []Datastream      `json:"Datastreams"`
	DatastreamsIotLink IotNavigationLink `json:"Datastreams@iot.navigationLink"`
}

// An ObservedProperty specifies the phenomenon of an Observation.
// See: https://docs.ogc.org/is/18-088/18-088.html#observedproperty
type ObservedProperty struct {
	Iot

	// A property provides a label for ObservedProperty entity, commonly a descriptive name.
	Name string `json:"name"`

	// The URI of the ObservedProperty. Dereferencing this URI SHOULD result in a representation of the definition of the ObservedProperty.
	Definition string `json:"definition"`

	// A description about the ObservedProperty.
	Description string `json:"description"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	Datastreams        []Datastream      `json:"Datastreams"`
	DatastreamsIotLink IotNavigationLink `json:"Datastreams@iot.navigationLink"`
}

// An Observation is the act of measuring or otherwise determining the value of a property [OGC 10-004r3 and ISO 19156:2011]
// See: https://docs.ogc.org/is/18-088/18-088.html#observation
type Observation struct {
	Iot

	// The time instant or period of when the Observation happens.
	PhenomenonTime time.Time `json:"phenomenonTime"`

	// The estimated value of an ObservedProperty from the Observation.
	Result Any `json:"result"`

	// The time of the Observation’s result was generated.
	ResultTime time.Time `json:"resultTime"`

	// Describes the quality of the result.
	ResultQuality string `json:"resultQuality"`

	// The time period during which the result may be used.
	ValidTime time.Time `json:"validTime"`

	// Key-value pairs showing the environmental conditions during measurement.
	Parameters struct {
		TargetPeriod  string `json:"targetPeriod"`
		MosmixElement string `json:"mosmix_element"`
	} `json:"parameters"`

	Datastream        *Datastream        `json:"Datastream"`
	FeatureOfInterest *FeatureOfInterest `json:"FeatureOfInterest"`

	DatastreamIotLink     IotNavigationLink `json:"Datastream@iot.navigationLink"`
	FeatureOfInterestLink IotNavigationLink `json:"FeatureOfInterest@iot.navigationLink"`
}

// An Observation results in a value being assigned to a phenomenon.
// The phenomenon is a property of a feature, the latter being the FeatureOfInterest of the Observation [OGC and ISO 19156:2011].
// In the context of the Internet of Things, many Observations’ FeatureOfInterest can be the Location of the Thing.
// For example, the FeatureOfInterest of a wifi-connect thermostat can be the Location of the thermostat (i.e., the living room where the thermostat is located in).
// In the case of remote sensing, the FeatureOfInterest can be the geographical area or volume that is being sensed.
// See: https://docs.ogc.org/is/18-088/18-088.html#featureofinterest
type FeatureOfInterest struct {
	Iot

	// A property provides a label for FeatureOfInterest entity, commonly a descriptive name.
	Name string `json:"name"`

	// The description about the FeatureOfInterest.
	Description string `json:"description"`

	// The encoding type of the feature property.
	EncodingType ValueCode `json:"encodingType"`

	// The detailed description of the feature. The data type is defined by encodingType.
	Feature Any `json:"feature"`

	// A JSON Object containing user-annotated properties as key-value pairs.
	Properties Properties `json:"properties"`

	Observations []Observation `json:"Observations"`

	ObservationsLink IotNavigationLink `json:"Observations@iot.navigationLink"`
}
