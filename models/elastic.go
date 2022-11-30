package models

type ElasticSearchSearch struct {
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
			ID int64 `json:"_id"`
			Score float32 `json:"_score"`
			Source *struct {
				Text string `json:"text"`
			}
		} `json:"hits"`
	}
}

// {
//     "_index": "article_text",
//     "_id": "24",
//     "_version": 1,
//     "_seq_no": 0,
//     "_primary_term": 1,
//     "found": true,
//     "_source": {
//         "text": "test2 text"
//     }
// }

type ElasticSearchGet struct {
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