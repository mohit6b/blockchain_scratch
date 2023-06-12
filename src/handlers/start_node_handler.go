package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
	"github.com/mohit6b/pow-blockchain-go/src/utils"
)

type StartNodeRequest struct {
	NodeID       string `json:"nodeID"`
	MinerAddress string `json:"minerAddress"`
}

type StartNodeResponse struct {
	Message string `json:"message"`
}

func HandleStartNode(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody StartNodeRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the miner address if provided
	if len(requestBody.MinerAddress) > 0 {
		if !utils.ValidateAddress(requestBody.MinerAddress) {
			http.Error(w, "Wrong miner address!", http.StatusBadRequest)
			return
		}
	}

	// Start the node server
	models.StartServer(requestBody.NodeID, requestBody.MinerAddress)

	// Create the response
	response := StartNodeResponse{
		Message: fmt.Sprintf("Node %s started successfully!", requestBody.NodeID),
	}

	// Marshal the response into JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the response body
	w.Write(responseJSON)
}
