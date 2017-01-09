package endpoints

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

func Cookie(collector bingo.EndpointCollector) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/cookies",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			data := map[string]map[string]string{
				"cookies": make(map[string]string),
			}

			for _, c := range r.Cookies() {
				data["cookies"][c.Name] = c.Value
			}

			WriteJSONResponse(w, data)
		},
	)
	endpoint.Description = "Returns cookie data."
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/cookies/set?name=value",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			var cookie *http.Cookie

			query := r.URL.Query()

			for k, _ := range query {
				cookie = &http.Cookie{
					Name: k,
					Value: query.Get(k),
					Path: "/",
				}

				http.SetCookie(w, cookie)
			}

			w.Header().Set("Location", "/cookies")
			w.WriteHeader(http.StatusFound)
		},
	)
	endpoint.Description = "Sets one or more simple cookies."
	endpoint.Query.Set("k1", "v1")
	endpoint.Query.Set("k2", "v2")
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/cookies/delete?name",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			var cookie *http.Cookie

			for k, _ := range r.URL.Query() {
				cookie = &http.Cookie{
					Name: k,
					Value: "",
					Path: "/",
					MaxAge: -1,
					Expires: time.Unix(0, 0),
				}

				http.SetCookie(w, cookie)
			}

			w.Header().Set("Location", "/cookies")
			w.WriteHeader(http.StatusFound)
		},
	)
	endpoint.Description = "Deletes one or more simple cookies."
	endpoint.Query.Set("k1", "")
	endpoint.Query.Set("k2", "")
	collector.AddEndpoint(endpoint)
}
