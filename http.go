package jsonds

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jbvmio/modules/httpserver"
	"github.com/julienschmidt/httprouter"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"
)

// GrafanaBackend a httpserver and version info.
type GrafanaBackend struct {
	APISrv     *httpserver.Module
	handlers   map[Endpoint]httprouter.Handle
	beHandlers map[Endpoint]BEHandler
}

// Endpoint represents a Datasource Endpoint Path.
type Endpoint string

// Main Endpoints defaults. Change according to your specific paths.
var (
	// RootEndpoint - (GET), used for status OK.
	RootEndpoint Endpoint = `/`

	// SearchEndpint - (POST), used for search.
	SearchEndpoint Endpoint = `/search`

	// QueryEndpoint - (POST), used for queries.
	QueryEndpoint Endpoint = `/query`

	// AnnotationsEndpoint - (POST), used for annotations.
	AnnotationsEndpoint Endpoint = `/annotations`

	// TagKeysEndpoint - (POST), used for tag keys (optional).
	TagKeysEndpoint Endpoint = `/tag-keys`

	// TagValuesEndpoint - (POST), used for tag values (optional).
	TagValuesEndpoint Endpoint = `/tag-values`
)

// Endpoints of the main JSON Datasource Grafana Backend.
var Endpoints = [...]*Endpoint{
	&RootEndpoint,
	&SearchEndpoint,
	&QueryEndpoint,
	&AnnotationsEndpoint,
	&TagKeysEndpoint,
	&TagValuesEndpoint,
}

// New configures httpserver and storage modules and returns a GrafanaBackend.
func New(config *Config) *GrafanaBackend {
	httpConfigs := httpserver.Configs{
		Server: make(map[string]*httpserver.Config, 1),
	}
	httpConfig := httpserver.NewConfig()
	httpConfig.Name = config.Name
	httpConfig.Address = config.HTTPAddress
	httpserver.LogLevel = config.LogLevel
	httpConfigs.Server[config.Name] = httpConfig
	apiSrv := httpserver.NewModule(&httpConfigs)

	return &GrafanaBackend{
		APISrv:     apiSrv,
		handlers:   make(map[Endpoint]httprouter.Handle, len(Endpoints)),
		beHandlers: make(map[Endpoint]BEHandler, len(Endpoints)),
	}
}

// Logger returns a child logger from the main GrafanaBackend server.
func (g *GrafanaBackend) Logger(name string) *zap.Logger {
	logger := g.APISrv.Logger.Named(name)
	return logger
}

// SetRoot configures the Root Endpoint Path.
func (g *GrafanaBackend) SetRoot(path string) {
	delete(g.beHandlers, RootEndpoint)
	RootEndpoint = Endpoint(path)
}

// SetSearch configures the Search Endpoint with the corresponding Path and BEHandler.
func (g *GrafanaBackend) SetSearch(path string, handler BEHandler) {
	delete(g.beHandlers, SearchEndpoint)
	SearchEndpoint = Endpoint(path)
	g.beHandlers[SearchEndpoint] = handler
}

// SetQuery configures the Query Endpoint with the corresponding Path and BEHandler.
func (g *GrafanaBackend) SetQuery(path string, handler BEHandler) {
	delete(g.beHandlers, QueryEndpoint)
	QueryEndpoint = Endpoint(path)
	g.beHandlers[QueryEndpoint] = handler
}

// SetAnnotations configures the Annotations Endpoint with the corresponding Path and BEHandler.
func (g *GrafanaBackend) SetAnnotations(path string, handler BEHandler) {
	delete(g.beHandlers, AnnotationsEndpoint)
	AnnotationsEndpoint = Endpoint(path)
	g.beHandlers[AnnotationsEndpoint] = handler
}

// SetTagKeys configures the TagKeys Endpoint with the corresponding Path and BEHandler.
func (g *GrafanaBackend) SetTagKeys(path string, handler BEHandler) {
	delete(g.beHandlers, TagKeysEndpoint)
	TagKeysEndpoint = Endpoint(path)
	g.beHandlers[TagKeysEndpoint] = handler
}

// SetTagValues configures the TagValues Endpoint with the corresponding Path and BEHandler.
func (g *GrafanaBackend) SetTagValues(path string, handler BEHandler) {
	delete(g.beHandlers, TagValuesEndpoint)
	TagValuesEndpoint = Endpoint(path)
	g.beHandlers[TagValuesEndpoint] = handler
}

// Configure sets all configurations.
func (g *GrafanaBackend) Configure() {
	defaultHandler := make(map[string]bool)
	for _, ep := range Endpoints {
		if *ep != RootEndpoint {
			valid, ok := g.beHandlers[*ep]
			switch {
			case !ok:
				g.beHandlers[*ep] = defaultBEHandler
				defaultHandler[string(*ep)] = true
			case valid == nil:
				g.beHandlers[*ep] = defaultBEHandler
				defaultHandler[string(*ep)] = true
			default:
				defaultHandler[string(*ep)] = false
			}
		}
	}
	g.APISrv.GET(string(RootEndpoint), g.statusOK)
	g.APISrv.POST(string(SearchEndpoint), g.handleSearch)
	g.APISrv.POST(string(QueryEndpoint), g.handleQuery)
	g.APISrv.POST(string(AnnotationsEndpoint), g.handleAnnotations)
	g.APISrv.POST(string(TagKeysEndpoint), g.handleTagKeys)
	g.APISrv.POST(string(TagValuesEndpoint), g.handleTagValues)
	g.APISrv.Configure()
	for _, ep := range Endpoints {
		if *ep != RootEndpoint {
			g.APISrv.Logger.Info("Configured Endpoint", zap.String("Path", string(*ep)), zap.Bool("default backend handler", defaultHandler[string(*ep)]))
		}
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if jsonBytes, err := json.Marshal(InvalidData{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":true,\"message\":\"could not encode JSON\",\"result\":{}}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(pretty.Pretty(jsonBytes))
	}
}

// Start starts the httpserver and storage modules.
func (g *GrafanaBackend) Start() {
	g.APISrv.Start()
}

// Stop stops the httpserver and storage modules.
func (g *GrafanaBackend) Stop() {
	err := g.APISrv.Stop()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
