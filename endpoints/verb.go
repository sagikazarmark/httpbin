package endpoints

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

func Verb(collector bingo.EndpointCollector) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/get",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			data := map[string]interface{}{
				"args":    Query(r.URL.Query()),
				"headers": Headers(r.Header),
				"origin":  ClientIp(r),
				"url":     r.URL.String(),
			}

			WriteJSONResponse(w, data)
		},
	)
	endpoint.Description = "Returns GET data."
	collector.AddEndpoint(endpoint)
}
