package endpoints

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

func Info(collector bingo.EndpointCollector) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/ip",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			data := map[string]string{
				"origin": ClientIp(r),
			}

			WriteJSONResponse(w, data)
		},
	)
	endpoint.Description = "Returns Origin IP."
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/user-agent",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			data := map[string]string{
				"user-agent": r.Header.Get("User-Agent"),
			}

			WriteJSONResponse(w, data)
		},
	)
	endpoint.Description = "Returns User-Agent."
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/headers",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			data := map[string]map[string]string{
				"headers": Headers(r.Header),
			}

			WriteJSONResponse(w, data)
		},
	)
	endpoint.Description = "Returns header dict."
	collector.AddEndpoint(endpoint)
}
