package jsonds

import "fmt"

// ResponseType identifies the type contained in the response.
type ResponseType string

// Available ResponseTypes:
const (
	RespAnnotation ResponseType = `annotation`
	RespTimeSeries ResponseType = `timeseries`
	RespTable      ResponseType = `table`
	RespTagKeys    ResponseType = `tagkeys`
	RespTagValues  ResponseType = `tagvalues`
	RespInvalid    ResponseType = `invalid`
)

// Response is the interface for the response to a Request.
type Response interface {
	// RespType should return the response type.
	RespType() ResponseType
	MarshalJSON() ([]byte, error)
}

// InvalidData is used to pass back a response that is invalid.
type InvalidData struct {
}

// RespType satisfies the QueryResponse interface and always returns the response type invalid.
func (t InvalidData) RespType() ResponseType {
	return RespInvalid
}

// MarshalJSON satisfies the QueryResponse interface and always an empty JSON response.
func (t InvalidData) MarshalJSON() ([]byte, error) {
	return []byte(`{}`), fmt.Errorf("could not handle request")
}

/*
// TimeSeries satisfies the QueryResponse interface and always returns an empty []TimeSeriesData.
func (t InvalidData) TimeSeries() []TimeSeriesData {
	return []TimeSeriesData{}
}

// Table satisfies the QueryResponse interface and always returns an empty []TableData.
func (t InvalidData) Table() []TableData {
	return []TableData{}
}
*/
