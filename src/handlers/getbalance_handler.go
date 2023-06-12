package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
	"github.com/mohit6b/pow-blockchain-go/src/utils"
)

type GetBalanceRequest struct {
	Address string `json:"address"`
	NodeID  string `json:"nodeID"`
}

type GetBalanceResponse struct {
	Balance int `json:"balance"`
}

func HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody GetBalanceRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the address
	if !utils.ValidateAddress(requestBody.Address) {
		http.Error(w, "ERROR: Address is not valid", http.StatusBadRequest)
		return
	}

	// Create the blockchain
	bc := models.NewBlockchain(requestBody.NodeID)
	defer bc.Db.Close()

	// Create the UTXO set
	UTXOSet := models.UTXOSet{bc}

	// Get the public key hash from the address
	pubKeyHash := utils.Base58Decode([]byte(requestBody.Address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	// Find the UTXOs for the address
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	// Calculate the balance
	balance := 0
	for _, out := range UTXOs {
		balance += out.Value
	}

	// Create the response
	response := GetBalanceResponse{
		Balance: balance,
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
