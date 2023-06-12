package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
	"github.com/mohit6b/pow-blockchain-go/src/utils"
)

type CreateBlockchainRequest struct {
	Address string `json:"address"`
	NodeID  string `json:"nodeID"`
}

type CreateBlockchainResponse struct {
	Message string `json:"message"`
}

func HandleCreateBlockchain(w http.ResponseWriter, r *http.Request) {
	print("handle create blockchain")
	// Parse the request body
	var requestBody CreateBlockchainRequest
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
	bc := models.CreateBlockchain(requestBody.Address, requestBody.NodeID)
	defer bc.Db.Close()

	// Create the UTXO set
	UTXOSet := models.UTXOSet{bc}
	UTXOSet.Reindex()

	// Create the response
	response := CreateBlockchainResponse{
		Message: "Blockchain created successfully!",
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
