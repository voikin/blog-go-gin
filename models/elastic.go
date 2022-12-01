package models

type ElasticSearchSearchResponse struct {
	Took int `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards *struct {
		Total int `json:"total"`
		Successful int `json:"successful"`
		Skipped int `json:"skipped"`
		Failed int `json:"failed"`
	} `json:"_shards"`
	Hits *struct {
		Total *struct{
			Value int `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float32 `json:"max_score"`
		Hits []*struct {
			Index string `json:"_index"`
			ID string `json:"_id"`
			Score float32 `json:"_score"`
			Source *struct {
				Text string `json:"text"`
			} `json:"_source"`
		} `json:"hits"`
	}
}

// {
//     "query": {
//         "ids": {
//             "values":[
//                 "40",
//                 "39"
//             ]
//         } 
//     }
// }

type ElasticSearchSearchPostRequest struct {
	Query *struct{
		IDs *struct{
			Values []int64 `json:"values"`
		} `json:"ids"`
	} `json:"query"`
}

type ElasticSearchGetResponse struct {
	Index string `json:"_index"`
	ID string `json:"_id"`
	Version int `json:"_version"`
	SeqNo int `json:"_seq_no"`
	PrimaryTerm int `json:"_primary_term"`
	Found bool `json:"found"`
	Source *struct{ 
		Text string `json:"text"`
	} `json:"_source"`
}