package jsonds

import "encoding/json"

// Datapoint are the metric values with unixtimestamp in milliseconds.
type Datapoint struct {
	MetricValue     float64
	UnixTimestampMS int64
}

// MarshalJSON provides JSON marshalling for a Datapoint.
func (d Datapoint) MarshalJSON() ([]byte, error) {
	var vals = [2]float64{
		d.MetricValue,
		float64(d.UnixTimestampMS),
	}
	return json.Marshal(vals)
}

// TimeSeriesData contains the datapoints for a TimeSeriesResponse.
type TimeSeriesData struct {
	Target     string      `json:"target"`
	Datapoints []Datapoint `json:"datapoints"`
}

// TimeSeriesResponse contains all the information needed to render a TimeSeries event.
type TimeSeriesResponse struct {
	Data []TimeSeriesData
}

// RespType satisfies the QueryResponse interface and returns the response type.
func (r TimeSeriesResponse) RespType() ResponseType {
	return RespTimeSeries
}

// MarshalJSON provides JSON marshalling for a TimeSeriesResponse.
func (r TimeSeriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}
