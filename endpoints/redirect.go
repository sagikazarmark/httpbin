package endpoints

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

func Redirect(collector bingo.EndpointCollector) {
	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/redirect/:n",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			redirectCount, err := strconv.ParseInt(p.ByName("n"), 10, 0)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			absolute := strings.ToLower(r.URL.Query().Get("absolute")) == "true"

			if 1 == redirectCount {
				if absolute {
					w.Header().Set("Location", AbsoluteRedirectURL("/get", r))
				} else {
					w.Header().Set("Location", "/get")
				}
			} else {
				if absolute {
					w.Header().Set("Location", "/absolute-redirect/"+strconv.FormatInt(redirectCount-1, 10))
				} else {
					w.Header().Set("Location", "/relative-redirect/"+strconv.FormatInt(redirectCount-1, 10))
				}
			}

			w.WriteHeader(http.StatusFound)
		},
	)
	endpoint.Description = "302 Redirects n times."
	endpoint.Parameters.Set("n", "6")
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/redirect-to?url=foo",
		func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			query := r.URL.Query()

			if _, ok := query["url"]; !ok {
				http.Error(w, "url parameter not passed", http.StatusBadRequest)
				return
			}

			w.Header().Set("Location", query.Get("url"))

			w.WriteHeader(http.StatusFound)
		},
	)
	endpoint.Description = "302 Redirects to the foo URL."
	endpoint.Query.Set("url", "/get")
	collector.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/relative-redirect/:n",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			redirectCount, err := strconv.ParseInt(p.ByName("n"), 10, 0)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if 1 == redirectCount {
				w.Header().Set("Location", "/get")
			} else {
				w.Header().Set("Location", "/relative-redirect/"+strconv.FormatInt(redirectCount-1, 10))
			}

			w.WriteHeader(http.StatusFound)
		},
	)
	endpoint.Description = "302 Relative redirects n times."
	endpoint.Parameters.Set("n", "6")
	collector.AddEndpoint(endpoint)


	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/absolute-redirect/:n",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			redirectCount, err := strconv.ParseInt(p.ByName("n"), 10, 0)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if 1 == redirectCount {
				w.Header().Set("Location", AbsoluteRedirectURL("/get", r))
			} else {
				w.Header().Set("Location", AbsoluteRedirectURL("/absolute-redirect/"+strconv.FormatInt(redirectCount-1, 10), r))
			}

			w.WriteHeader(http.StatusFound)
		},
	)
	endpoint.Description = "302 Absolute redirects n times."
	endpoint.Parameters.Set("n", "6")
	collector.AddEndpoint(endpoint)
}

func AbsoluteRedirectURL(path string, r *http.Request) (string) {
	url := *r.URL

	url.Path = path
	url.RawQuery = ""
	url.Fragment = ""

	return url.String()
}
