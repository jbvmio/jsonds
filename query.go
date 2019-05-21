package jsonds

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
