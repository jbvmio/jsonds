package jsonds

import (
	"encoding/json"
	"time"
)

// AnnotationsReq encodes the information provided by Grafana in its requests.
type AnnotationsReq struct {
	Range      Range      `json:"range"`
	Annotation Annotation `json:"annotation"`
}

// Range specifies the time range the request is valid for.
type Range struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// Annotation is the object passed by Grafana when it fetches annotations.
//
// http://docs.grafana.org/plugins/developing/datasources/#annotation-query
type Annotation struct {
	// Name must match in the request and response
	Name string `json:"name"`

	Datasource string `json:"datasource"`
	IconColor  string `json:"iconColor"`
	Enable     bool   `json:"enable"`
	ShowLine   bool   `json:"showLine"`
	Query      string `json:"query"`
}

// AnnotationResponse contains all the information needed to render an
// annotation event.
//
// https://github.com/grafana/simple-json-datasource#annotation-api
type AnnotationResponse struct {
	// The original annotation sent from Grafana.
	Annotation Annotation `json:"annotation"`
	// Time since UNIX Epoch in milliseconds. (required)
	Time int64 `json:"time"`
	// The title for the annotation tooltip. (required)
	Title string `json:"title"`
	// Tags for the annotation. (optional)
	Tags string `json:"tags"`
	// Text for the annotation. (optional)
	Text string `json:"text"`
}

// RespType satisfies the QueryResponse interface and returns the response type.
func (r AnnotationResponse) RespType() ResponseType {
	return RespAnnotation
}

// MarshalJSON provides JSON marshalling for a TimeSeriesResponse.
func (r AnnotationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

// AnnotationQuery is a collection of possible filters for a Grafana annotation
// request.
//
// These values are set in gear icon > annotations > edit > Query and can have
// such useful values as:
//
//		{"example1": "$example1", "example2": "$example2"}
//
// or if you do not want filtering leave it blank.
type AnnotationQuery struct {
	Example1 string `json:"example1"`
	Example2 string `json:"example2"`
}

// ReqType returns the Request type.
func (r *AnnotationsReq) ReqType() RequestType {
	return ReqAnnotation
}

// Search returns the SearchRequest.
func (r *AnnotationsReq) Search() *SearchRequest {
	return nil
}

// Anno returns the AnnotationsReq
func (r *AnnotationsReq) Anno() *AnnotationsReq {
	return r
}

// Query returns the QueryRequest
func (r *AnnotationsReq) Query() *QueryRequest {
	return nil
}

// TagKeys returns the AnnotationsReq
func (r *AnnotationsReq) TagKeys() *TagKeysReq {
	return nil
}

// TagValues returns the AnnotationsReq
func (r *AnnotationsReq) TagValues() *TagValuesReq {
	return nil
}
