package jsonds

import "encoding/json"

// KeyType identifies the type for a Key.
type KeyType string

// KeyTypes for available TagKeys:
const (
	KeyTypeString KeyType = `string`
	KeyTypeNumber KeyType = `number`
	KeyTypeTime   KeyType = `time`
)

// KeyTypes array of all available types:
var KeyTypes = [...]KeyType{
	KeyTypeString,
	KeyTypeNumber,
	KeyTypeTime,
}

// TagKeysReq describes a TagKeys Request.
type TagKeysReq map[string]interface{}

// TagKey describes a TagKey.
type TagKey struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// TagValue contains the text value for a TagKey.
type TagValue struct {
	Text string `json:"text"`
}

// TagPair maps a TagKey with a TagValue.
type TagPair map[TagKey]TagValue

// TagValuesReq describes a TagValues Request.
type TagValuesReq struct {
	Key string `json:"key"`
}

// TagKeysResp contains the response for a TagsValuesReq.
type TagKeysResp struct {
	Data []TagKey
}

// RespType satisfies the QueryResponse interface and returns the response type.
func (r TagKeysResp) RespType() ResponseType {
	return RespTagValues
}

// MarshalJSON provides JSON marshalling for a TimeSeriesResponse.
func (r TagKeysResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

// TagValuesResp contains the response for a TagsValuesReq.
type TagValuesResp struct {
	Data []TagValue
}

// RespType satisfies the QueryResponse interface and returns the response type.
func (r TagValuesResp) RespType() ResponseType {
	return RespTagValues
}

// MarshalJSON provides JSON marshalling for a TimeSeriesResponse.
func (r TagValuesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

// ReqType returns the Request type.
func (r *TagKeysReq) ReqType() RequestType {
	return ReqTagKeys
}

// Search returns the SearchRequest.
func (r *TagKeysReq) Search() *SearchRequest {
	return nil
}

// Anno returns the AnnotationsReq
func (r *TagKeysReq) Anno() *AnnotationsReq {
	return nil
}

// Query returns the QueryRequest
func (r *TagKeysReq) Query() *QueryRequest {
	return nil
}

// TagKeys returns the AnnotationsReq
func (r *TagKeysReq) TagKeys() *TagKeysReq {
	return r
}

// TagValues returns the AnnotationsReq
func (r *TagKeysReq) TagValues() *TagValuesReq {
	return nil
}

// ReqType returns the Request type.
func (r *TagValuesReq) ReqType() RequestType {
	return ReqTagValue
}

// Search returns the SearchRequest.
func (r *TagValuesReq) Search() *SearchRequest {
	return nil
}

// Anno returns the AnnotationsReq
func (r *TagValuesReq) Anno() *AnnotationsReq {
	return nil
}

// Query returns the QueryRequest
func (r *TagValuesReq) Query() *QueryRequest {
	return nil
}

// TagKeys returns the AnnotationsReq
func (r *TagValuesReq) TagKeys() *TagKeysReq {
	return nil
}

// TagValues returns the AnnotationsReq
func (r *TagValuesReq) TagValues() *TagValuesReq {
	return r
}
