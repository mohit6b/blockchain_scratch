package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
)

type ReindexUTXORequest struct {
	NodeID string `json:"nodeID"`
}

type ReindexUTXOResponse struct {
	TransactionCount int `json:"transactionCount"`
}

func HandleReindexUTXO(w http.ResponseWriter, r *http.Request) {
	var requestBody ReindexUTXORequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	bc := models.NewBlockchain(requestBody.NodeID)
	UTXOSet := models.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
	// Create the response
	response := ReindexUTXOResponse{
		TransactionCount: count,
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
