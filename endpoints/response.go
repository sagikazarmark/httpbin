package endpoints

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

func Response(collector bingo.EndpointCollector) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/status/:code",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			rawCode := p.ByName("code")

			if strings.Contains(rawCode, ",") {
				codes := strings.Split(rawCode, ",")
				rawCode = strings.TrimSpace(codes[rand.Intn(len(codes))])
			}

			code, err := strconv.ParseInt(rawCode, 10, 0)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(int(code))

			if 418 == code {
				teapot, err := Asset("data/templates/teapot")
				if err != nil {
					panic(err)
				}

				w.Write(teapot)
			}
		},
	)
	endpoint.Description = "Returns given HTTP Status code."
	endpoint.Parameters.Set("code", "418")
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/response-headers?key=val",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

			for k, v := range r.URL.Query() {
				for _, header := range v {
					w.Header().Add(k, header)
				}
			}

			data := Headers(w.Header())

			WriteJSONResponse(w, data)
		},
	)
	endpoint.Description = "Returns given response headers."
	endpoint.Query.Set("Server", "httpbin")
	collector.AddEndpoint(endpoint)
}
