package jsonds

import "encoding/json"

// SearchRequest is used for parsing Grafana requests for variable names.
type SearchRequest struct {
	Target string `json:"target"`
}

// SearchResponse contains the values returned from a SearchRequest.
type SearchResponse struct {
	Data []string
}

// RespType satisfies the QueryResponse interface and returns the response type.
func (r SearchResponse) RespType() ResponseType {
	return RespTable
}

// MarshalJSON provides JSON marshalling for a TimeSeriesResponse.
func (r SearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

// ReqType returns the Request type.
func (r *SearchRequest) ReqType() RequestType {
	return ReqSearch
}

// Search returns the SearchRequest.
func (r *SearchRequest) Search() *SearchRequest {
	return r
}

// Anno returns the AnnotationsReq
func (r *SearchRequest) Anno() *AnnotationsReq {
	return nil
}

// Query returns the QueryRequest
func (r *SearchRequest) Query() *QueryRequest {
	return nil
}

// TagKeys returns the AnnotationsReq
func (r *SearchRequest) TagKeys() *TagKeysReq {
	return nil
}

// TagValues returns the AnnotationsReq
func (r *SearchRequest) TagValues() *TagValuesReq {
	return nil
}
