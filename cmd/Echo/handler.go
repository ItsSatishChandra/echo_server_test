package echo

import (
	http "net/http"
)

func HandleEcho(w http.ResponseWriter, r *http.Request) {

	input := ""
	switch r.Method {
	case http.MethodGet:
		input := r.URL.Query().Get("input")
		w.Write([]byte(input))
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		input = r.FormValue("input")

	default:
		input = "Method not allowed"
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	Echo(input)
}

func HandleGet(w http.ResponseWriter, r *http.Request) string {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return "No input received, only query Parameter input is allowed"
	}
	input := r.URL.Query().Get("input")
	return input
}
