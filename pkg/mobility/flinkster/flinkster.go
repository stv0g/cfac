package flinkster

// https://data.deutschebahn.com/dataset/flinkster-api.html

const (
	UrlApi              = "http://api.deutschebahn.com/flinkster-api-ng/v1"
	UrlApiAreas         = UrlApi + "/areas?lat={lat}&lon={lon}&radius={radius}}&providernetwork={network}"
	UrlBookingProposals = UrlApi + "/bookingproposals?lat={lat}&log={lon}&radius={radius}&providernetwork={network}"
)
