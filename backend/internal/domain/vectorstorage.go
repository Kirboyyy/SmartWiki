package domain

type SearchRequest struct {
	Query          string
	MetaDataFilter map[string]interface{}
	TopK           int
}

type SearchResult struct {
	Score    float32
	MetaData map[string]interface{}
	Document Page
}
