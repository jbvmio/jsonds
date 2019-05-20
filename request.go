package jsonds

// RequestType identifies the target Endpoint in the request.
type RequestType string

// Available RequestType:
const (
	ReqAnnotation RequestType = `annotation`
	ReqSearch     RequestType = `search`
	ReqQuery      RequestType = `query`
	ReqTagKeys    RequestType = `tagkeys`
	ReqTagValue   RequestType = `tagvalues`
	ReqInvalid    RequestType = `invalid`
)

// Request is the interface for Backend Request.
// A Request should return one of the available RequestTypes.
type Request interface {
	// ReqType should return the Request type.
	ReqType() RequestType

	// Search returns the SeaerchRequest
	Search() *SearchRequest

	// Anno returns the AnnotationsReq
	Anno() *AnnotationsReq

	// Query returns the QueryRequest
	Query() *QueryRequest

	// TagKeys returns the TagKeysReq
	TagKeys() *TagKeysReq

	// TagValues returns the TagValuesReq
	TagValues() *TagValuesReq
}
