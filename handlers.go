package jsonds

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"
)

// BEHandler (BackEnd Handler) functions handle Query Requests and return a QueryResponse and error.
type BEHandler func(Request) (Response, error)

// defaultBEHandler is a default EPHandler which passes an invalid response.
func defaultBEHandler(req Request) (Response, error) {
	var ep Endpoint
	switch req.ReqType() {
	case ReqAnnotation:
		ep = AnnotationsEndpoint
	case ReqSearch:
		ep = SearchEndpoint
	case ReqQuery:
		ep = QueryEndpoint
	case ReqTagKeys:
		ep = TagKeysEndpoint
	case ReqTagValue:
		ep = TagValuesEndpoint
	default:
		ep = `unknown path`
	}
	return InvalidData{}, fmt.Errorf("default endpoint handler: could not handle request for %v", ep)
}

func (g *GrafanaBackend) statusOK(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g.APISrv.Logger.Debug("root endpoint called", zap.String("method", r.Method), zap.String("from", r.RemoteAddr), zap.String("URI", r.RequestURI))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (g *GrafanaBackend) handleSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g.APISrv.Logger.Debug("search endpoint called", zap.String("endpoint", string(SearchEndpoint)), zap.String("method", r.Method), zap.String("from", r.RemoteAddr), zap.String("URI", r.RequestURI))
	switch r.Method {
	case http.MethodPost:
		var req SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("json decode failure", zap.Error(err))
			errMsg := fmt.Sprintf("json decode failure: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}
		resp, err := g.beHandlers[SearchEndpoint](&req)
		if err != nil {
			logger.Error("backend handler failer", zap.String("endpoint", string(SearchEndpoint)), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{\"error\":true,\"message\":\"` + err.Error() + `\"}`))
		}
		g.writeJSONResponse(w, http.StatusOK, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad method; supported POST"))
		return
	}
}

func (g *GrafanaBackend) handleQuery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g.APISrv.Logger.Debug("query endpoint called", zap.String("endpoint", string(QueryEndpoint)), zap.String("method", r.Method), zap.String("from", r.RemoteAddr), zap.String("URI", r.RequestURI))
	switch r.Method {
	case http.MethodPost:
		var req QueryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("json decode failure", zap.Error(err))
			errMsg := fmt.Sprintf("json decode failure: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}
		resp, err := g.beHandlers[QueryEndpoint](&req)
		if err != nil {
			logger.Error("backend handler failer", zap.String("endpoint", string(QueryEndpoint)), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{\"error\":true,\"message\":\"` + err.Error() + `\"}`))
		}
		g.writeJSONResponse(w, http.StatusOK, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad method; supported POST"))
		return
	}
}

func (g *GrafanaBackend) handleAnnotations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g.APISrv.Logger.Debug("annotation endpoint called", zap.String("endpoint", string(AnnotationsEndpoint)), zap.String("method", r.Method), zap.String("from", r.RemoteAddr), zap.String("URI", r.RequestURI))
	switch r.Method {
	case http.MethodPost:
		var req AnnotationsReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("json decode failure", zap.Error(err))
			errMsg := fmt.Sprintf("json decode failure: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}
		resp, err := g.beHandlers[AnnotationsEndpoint](&req)
		if err != nil {
			logger.Error("backend handler failer", zap.String("endpoint", string(AnnotationsEndpoint)), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{\"error\":true,\"message\":\"` + err.Error() + `\"}`))
		}
		g.writeJSONResponse(w, http.StatusOK, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad method; supported POST"))
		return
	}
}

func (g *GrafanaBackend) handleTagKeys(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g.APISrv.Logger.Debug("tagkeys endpoint called", zap.String("endpoint", string(TagKeysEndpoint)), zap.String("method", r.Method), zap.String("from", r.RemoteAddr), zap.String("URI", r.RequestURI))
	switch r.Method {
	case http.MethodPost:
		var req TagKeysReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("json decode failure", zap.Error(err))
			errMsg := fmt.Sprintf("json decode failure: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}
		resp, err := g.beHandlers[TagKeysEndpoint](&req)
		if err != nil {
			logger.Error("backend handler failer", zap.String("endpoint", string(TagKeysEndpoint)), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{\"error\":true,\"message\":\"` + err.Error() + `\"}`))
		}
		g.writeJSONResponse(w, http.StatusOK, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad method; supported POST"))
		return
	}
}

func (g *GrafanaBackend) handleTagValues(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g.APISrv.Logger.Debug("tagvalues endpoint called", zap.String("endpoint", string(TagValuesEndpoint)), zap.String("method", r.Method), zap.String("from", r.RemoteAddr), zap.String("URI", r.RequestURI))
	switch r.Method {
	case http.MethodPost:
		var req TagValuesReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("json decode failure", zap.Error(err))
			errMsg := fmt.Sprintf("json decode failure: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}
		resp, err := g.beHandlers[TagValuesEndpoint](&req)
		if err != nil {
			logger.Error("backend handler failer", zap.String("endpoint", string(TagValuesEndpoint)), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{\"error\":true,\"message\":\"` + err.Error() + `\"}`))
		}
		g.writeJSONResponse(w, http.StatusOK, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad method; supported POST"))
		return
	}
}

// WriteJSONResponse generates a JSON response from the given JSON object and writes to the given ResponseWriter.
func (g *GrafanaBackend) writeJSONResponse(w http.ResponseWriter, statusCode int, jsonObj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if jsonBytes, err := json.Marshal(jsonObj); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":true,\"message\":\"could not encode JSON\",\"result\":{}}"))
	} else {
		w.WriteHeader(statusCode)
		w.Write(pretty.Pretty(jsonBytes))
	}
}

/*
type httpResponseRequestInfo struct {
	URI  string `json:"url"`
	Host string `json:"host"`
}

type httpResponseError struct {
	Error   bool                    `json:"error"`
	Message string                  `json:"message"`
	Request httpResponseRequestInfo `json:"request"`
}

type listResponse struct {
	Request       httpResponseRequestInfo `json:"request"`
	ListType      string                  `json:"listType"`
	TopicsChecked map[string][]string     `json:"topicsChecked"`
}

type parityResponse struct {
	Request  httpResponseRequestInfo `json:"request"`
	ListType string                  `json:"listType"`
	//Current  ParityCounts            `json:"current"`
}

func makeRequestInfo(r *http.Request) httpResponseRequestInfo {
	hostname, _ := os.Hostname()
	return httpResponseRequestInfo{
		URI:  r.URL.Path,
		Host: hostname,
	}
}
*/
