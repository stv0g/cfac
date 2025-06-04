// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package noq

const (
	// Seems to be some sort of Elasticserach query
	Url = "https://es-data-proxy.no-q.info/prod/" // + "?source=<src>&source_content_type=application%2Fjson"

	// <src> is URLencoded: {"size":0,"query":{"bool":{"filter":[{"terms":{"tags":["aachen"]}},{"bool":{"should":[{"bool":{"filter":[{"range":{"checkin_at":{"gte":"now","lt":"2021-11-15T23:59:59.999+01:00"}}},{"range":{"min_booking_date":{"lte":"now"}}}]}},{"term":{"placeholder":true}}]}}]}},"aggs":{"gyms":{"terms":{"field":"gym_id","size":10000},"aggs":{"gym":{"top_hits":{"_source":"*","size":1}},"total_free_spots":{"sum":{"field":"free_spots"}},"placeholder":{"terms":{"field":"placeholder"}}}}}}
)
