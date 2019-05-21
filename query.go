package jsonds

import (
	"strings"

	"github.com/spf13/cast"
)

// QueryRequest encodes the information provided by Grafana in its requests.
type QueryRequest struct {
	Range         Range         `json:"range"`
	Interval      string        `json:"interval"`
	InvervalMS    int64         `json:"intervalMs"`
	Targets       []Target      `json:"targets"`
	AdhocFilters  []AdhocFilter `json:"adhocFilters"`
	Format        string        `json:"format"`
	MaxDataPoints int           `json:"maxDataPoints"`
	ScopedVars    ScopedVar     `json:"scopedVars"`
}

// GetGlobalVar returns ScopedPaired Variables by the given variable name.
func (r *QueryRequest) GetGlobalVar(variable string) ScopedPair {
	return r.ScopedVars[variable]
}

// ListGlobalVars returns a list of all ScopedPaired Variables by variable name.
func (r *QueryRequest) ListGlobalVars() []string {
	var variables []string
	for v := range r.ScopedVars {
		variables = append(variables, v)
	}
	return variables
}

// ListTargets returns a list of all Targets.
func (r *QueryRequest) ListTargets() []string {
	var targets []string
	for _, t := range r.Targets {
		targets = append(targets, t.Target)
	}
	return targets
}

// Target specifies the intended target of a request.
type Target struct {
	Target string                 `json:"target"`
	Type   string                 `json:"type"`
	Data   map[string]interface{} `json:"data"`
}

// GetVar returns Variables by the given variable name.
func (t *Target) GetVar(variable string) interface{} {
	return t.Data[variable]
}

// GetVarStrings parses and returns Variables by the given variable name as a string array.
func (t *Target) GetVarStrings(variable string) []string {
	var strArray []string
	switch i := t.Data[variable].(type) {
	case []string:
		strArray = i
	case string:
		strArray = strings.Split(strings.Trim(i, `{}`), `,`)
	case []interface{}:
		for _, x := range flattenArray(i) {
			strArray = append(strArray, toStringArray(x)...)
		}
	default:
		strArray = toStringArray(i)
	}
	return strArray
}

// ListVars returns a list of all Variables by variable name.
func (t *Target) ListVars() []string {
	var variables []string
	for v := range t.Data {
		variables = append(variables, v)
	}
	return variables
}

// AdhocFilter holds adhoc key values.
type AdhocFilter struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// ScopedPair contains a ScopedVar pair.
type ScopedPair struct {
	Text  interface{} `json:"text"`
	Value interface{} `json:"value"`
}

// TextStrings parses and if possible returns the Text of a ScopedPair as a string array.
func (s ScopedPair) TextStrings() []string {
	return toStringArray(s.Text)
}

// ValueStrings parses and if possible returns the values of a ScopedPair as a string array.
func (s ScopedPair) ValueStrings() []string {
	return toStringArray(s.Value)
}

func toStringArray(s interface{}) []string {
	var strArray []string
	switch i := s.(type) {
	case []interface{}:
		vals := flattenArray(i)
		for _, v := range vals {
			strArray = append(strArray, cast.ToString(v))
		}
	case []int:
		for _, v := range i {
			strArray = append(strArray, cast.ToString(v))
		}
	case []int64:
		for _, v := range i {
			strArray = append(strArray, cast.ToString(v))
		}
	case []float32:
		for _, v := range i {
			strArray = append(strArray, cast.ToString(v))
		}
	case []float64:
		for _, v := range i {
			strArray = append(strArray, cast.ToString(v))
		}
	case []string:
		for _, v := range i {
			strArray = append(strArray, v)
		}
	case string:
		strArray = append(strArray, i)
	}
	return strArray
}

func flattenArray(i []interface{}) []interface{} {
	var returnArray []interface{}
	var workArray [][]interface{}
	workArray[0] = i
	for x := 0; x < len(workArray); x++ {
		for y := 0; y < len(workArray[x]); y++ {
			switch z := workArray[x][y].(type) {
			case []interface{}:
				workArray = append(workArray, z)
			case interface{}:
				returnArray = append(returnArray, z)
			case int, int64, float32, float64, string:
				returnArray = append(returnArray, z)
			case []int:
				for _, n := range z {
					returnArray = append(returnArray, n)
				}
			case []int64:
				for _, n := range z {
					returnArray = append(returnArray, n)
				}
			case []float32:
				for _, n := range z {
					returnArray = append(returnArray, n)
				}
			case []float64:
				for _, n := range z {
					returnArray = append(returnArray, n)
				}
			case []string:
				for _, n := range z {
					returnArray = append(returnArray, n)
				}
			}
		}
	}
	return returnArray
}

// ScopedVar contains ScopedVariable Pairs.
type ScopedVar map[string]ScopedPair

// ReqType returns the Request type.
func (r *QueryRequest) ReqType() RequestType {
	return ReqQuery
}

// Search returns the SearchRequest.
func (r *QueryRequest) Search() *SearchRequest {
	return nil
}

// Anno returns the AnnotationsReq
func (r *QueryRequest) Anno() *AnnotationsReq {
	return nil
}

// Query returns the QueryRequest
func (r *QueryRequest) Query() *QueryRequest {
	return r
}

// TagKeys returns the AnnotationsReq
func (r *QueryRequest) TagKeys() *TagKeysReq {
	return nil
}

// TagValues returns the AnnotationsReq
func (r *QueryRequest) TagValues() *TagValuesReq {
	return nil
}
