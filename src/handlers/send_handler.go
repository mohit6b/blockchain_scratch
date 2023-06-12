package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohit6b/pow-blockchain-go/src/models"
	"github.com/mohit6b/pow-blockchain-go/src/utils"
)

var knownNodes = []string{"localhost:3000"}

type SendRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Amount  int    `json:"amount"`
	NodeID  string `json:"nodeID"`
	MineNow bool   `json:"mineNow"`
}

type SendResponse struct {
	Message string `json:"message"`
}

func HandleSend(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody SendRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate sender and recipient addresses
	if !utils.ValidateAddress(requestBody.From) {
		http.Error(w, "ERROR: Sender address is not valid", http.StatusBadRequest)
		return
	}
	if !utils.ValidateAddress(requestBody.To) {
		http.Error(w, "ERROR: Recipient address is not valid", http.StatusBadRequest)
		return
	}

	// Create the blockchain
	bc := models.NewBlockchain(requestBody.NodeID)
	defer bc.Db.Close()

	// Create the UTXO set
	UTXOSet := models.UTXOSet{bc}

	// Create the wallets
	wallets, err := models.NewWallets(requestBody.NodeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the sender's wallet
	wallet := wallets.GetWallet(requestBody.From)

	// Create the transaction
	tx := models.NewUTXOTransaction(&wallet, requestBody.To, requestBody.Amount, &UTXOSet)

	// Handle the transaction based on mineNow flag
	if requestBody.MineNow {
		// Create the coinbase transaction
		cbTx := models.NewCoinbaseTX(requestBody.From, "")
		txs := []*models.Transaction{cbTx, tx}

		// Mine a new block
		newBlock := bc.MineBlock(txs)

		// Update the UTXO set
		UTXOSet.Update(newBlock)
	} else {
		models.SendTx(knownNodes[0], tx)
	}

	// Create the response
	response := SendResponse{
		Message: "Success!",
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
