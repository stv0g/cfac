package noq

type AutoGenerated struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   Shard `json:"_shards"`
	Hits Hits `json:"hits"`
	Aggregations struct {
		Gyms Gyms `json:"gyms"`
	} `json:"aggregations"`
}

type Shard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Gyms struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []Bucket `json:"buckets"`
}
type Gym struct {
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []Hit `json:"hits"`
	} `json:"hits"`
}

type Bucket struct {
	Key            int `json:"key"`
	DocCount       int `json:"doc_count"`
	TotalFreeSpots struct {
		Value float64 `json:"value"`
	} `json:"total_free_spots"`
	Gym Gym `json:"gym"`
	Placeholder Placeholder `json:"placeholder"`
}

type Placeholder struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key         int    `json:"key"`
		KeyAsString string `json:"key_as_string"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

type Source struct {
	GymID           int         `json:"gym_id"`
	Zip             string      `json:"zip"`
	CheckinAt       time.Time   `json:"checkin_at"`
	Country         string      `json:"country"`
	Website         interface{} `json:"website"`
	Address         string      `json:"address"`
	Lng             float64     `json:"lng"`
	City            string      `json:"city"`
	SlimTestbuchen  bool        `json:"slim_testbuchen"`
	Tags            []string    `json:"tags"`
	FreeSpots       int         `json:"free_spots"`
	AppIntegrations []string    `json:"app_integrations"`
	Name            string      `json:"name"`
	MinBookingDate  time.Time   `json:"min_booking_date"`
	ID              string      `json:"id"`
	Placeholder     bool        `json:"placeholder"`
	Prices          struct {} `json:"prices"`
	Slug string  `json:"slug"`
	Lat  float64 `json:"lat"`
	Hash string  `json:"hash"`
}

type Hits struct {
	Total struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore interface{}   `json:"max_score"`
	Hits     []interface{} `json:"hits"`
}

type Hit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source `json:"_source"`
}


type RequestSource struct {
	Size  int `json:"size"`
	Query Query `json:"query"`
	Aggs struct {
		Gyms struct {
			Terms struct {
				Field string `json:"field"`
				Size  int    `json:"size"`
			} `json:"terms"`
			Aggs struct {
				Gym struct {
					TopHits struct {
						Source string `json:"_source"`
						Size   int    `json:"size"`
					} `json:"top_hits"`
				} `json:"gym"`
				TotalFreeSpots struct {
					Sum struct {
						Field string `json:"field"`
					} `json:"sum"`
				} `json:"total_free_spots"`
				Placeholder struct {
					Terms struct {
						Field string `json:"field"`
					} `json:"terms"`
				} `json:"placeholder"`
			} `json:"aggs"`
		} `json:"gyms"`
	} `json:"aggs"`
}

type Query struct {
	Bool struct {
		Filter []struct {
			Terms struct {
				Tags []string `json:"tags"`
			} `json:"terms,omitempty"`
			Bool struct {
				Should []struct {
					Bool struct {
						Filter []struct {
							Range struct {
								CheckinAt struct {
									Gte string    `json:"gte"`
									Lt  time.Time `json:"lt"`
								} `json:"checkin_at"`
							} `json:"range"`
						} `json:"filter"`
					} `json:"bool,omitempty"`
					Term struct {
						Placeholder bool `json:"placeholder"`
					} `json:"term,omitempty"`
				} `json:"should"`
			} `json:"bool,omitempty"`
		} `json:"filter"`
	} `json:"bool"`
}