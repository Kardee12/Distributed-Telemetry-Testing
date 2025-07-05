package handlers

import (
	"log"
	"net/http"
	"telem.kmani/internal/api/handlers/helpers"
	"telem.kmani/internal/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	client := utils.GetClient()

	res, err := client.Ping()
	if err != nil || res.StatusCode != 200 {
		helpers.WriteError(w, http.StatusServiceUnavailable, "OpenSearch is not available")
		return
	}
	defer res.Body.Close()
	log.Println("HealthCheck Complete")
	helpers.WriteJSON(w, http.StatusOK, helpers.NewAPIResponse("PING"))
}
