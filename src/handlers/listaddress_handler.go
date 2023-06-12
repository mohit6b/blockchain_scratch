package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
)

type ListAddressesRequest struct {
	NodeID string `json:"nodeID"`
}

type ListAddressesResponse struct {
	Addresses []string `json:"addresses"`
}

func HandleListAddress(w http.ResponseWriter, r *http.Request) {
	var requestBody ListAddressesRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	wallets, err := models.NewWallets(requestBody.NodeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	addresses := wallets.GetAddresses()

	// Create the response
	response := ListAddressesResponse{
		Addresses: addresses,
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
