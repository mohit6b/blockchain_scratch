package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
)

type CreateWalletRequest struct {
	NodeID string `json:"nodeID"`
}

type CreateWalletResponse struct {
	Address string `json:"address"`
}

func HandleCreateWallet(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody CreateWalletRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Create the wallets
	wallets, _ := models.NewWallets(requestBody.NodeID)

	// Create a new wallet and get the address
	address := wallets.CreateWallet()

	// Save the wallets to file
	wallets.SaveToFile(requestBody.NodeID)

	// Create the response
	response := CreateWalletResponse{
		Address: address,
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
