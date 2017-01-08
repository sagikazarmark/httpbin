package endpoints

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func ClientIp(r *http.Request) string {
	if _, ok := r.Header["X-Forwarded-For"]; ok {
		return r.Header.Get("X-Forwarded-For")
	}

	return r.RemoteAddr
}

func Headers(h http.Header) map[string]string {
	headers := make(map[string]string)

	for k, _ := range h {
		headers[k] = h.Get(k)
	}

	return headers
}

func Query(q url.Values) map[string]interface{} {
	query := make(map[string]interface{})

	for k, v := range q {
		if len(v) == 1 {
			query[k] = v[0]
		} else {
			query[k] = v
		}
	}

	return query
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
