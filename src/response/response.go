package response

import (
	"encoding/json"
	"net/http"
)

type Data map[string]interface{}

type Response struct {
	Message string
	Status  int
}

func JsonResponse(w http.ResponseWriter, status int, data Data) {
	data["status"] = status

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
