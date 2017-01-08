package endpoints

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

func Encoding(collector bingo.EndpointCollector) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/encoding/utf8",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			data, err := Asset("data/templates/utf8")
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(data)
		},
	)
	endpoint.Description = "Returns page containing UTF-8 data."
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/gzip",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			gz := gzip.NewWriter(w)
			defer gz.Close()

			w.Header().Set("Content-Encoding", "gzip")

			data := map[string]interface{}{
				"gzipped": true,
				"headers": Headers(r.Header),
				"method":  r.Method,
				"origin":  ClientIp(r),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(gz).Encode(data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
	endpoint.Description = "Returns gzip-encoded data."
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/deflate",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			df, _ := flate.NewWriter(w, 1)
			defer df.Close()

			w.Header().Set("Content-Encoding", "deflate")

			data := map[string]interface{}{
				"deflated": true,
				"headers":  Headers(r.Header),
				"method":   r.Method,
				"origin":   ClientIp(r),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(df).Encode(data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
	endpoint.Description = "Returns deflate-encoded data."
	collector.AddEndpoint(endpoint)
}
