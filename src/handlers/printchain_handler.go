package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mohit6b/pow-blockchain-go/src/models"
)

type PrintChainRequest struct {
	NodeID string `json:"nodeID"`
}

type BlockData struct {
	Hash          string   `json:"hash"`
	Height        int      `json:"height"`
	PrevBlockHash string   `json:"prevBlockHash"`
	ProofOfWork   string   `json:"proofOfWork"`
	Transactions  []string `json:"transactions"`
}

type PrintChainResponse struct {
	Blocks []BlockData `json:"blocks"`
}

func HandlePrintChain(w http.ResponseWriter, r *http.Request) {
	// Retrieve the nodeID from the request URL or query parameters
	// Parse the request body
	var requestBody PrintChainRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	bc := models.NewBlockchain(requestBody.NodeID)
	defer bc.Db.Close()

	bci := bc.Iterator()

	// Initialize the response blocks slice
	blocks := []BlockData{}

	// Iterate over the blockchain and collect block data
	for {
		block := bci.Next()

		blockData := BlockData{
			Hash:          fmt.Sprintf("%x", block.Hash),
			Height:        block.Height,
			PrevBlockHash: fmt.Sprintf("%x", block.PrevBlockHash),
			ProofOfWork:   strconv.FormatBool(models.NewProofOfWork(block).Validate()),
			Transactions:  []string{},
		}

		for _, tx := range block.Transactions {
			blockData.Transactions = append(blockData.Transactions, fmt.Sprintf("%v", tx))
		}

		blocks = append(blocks, blockData)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	// Create the response
	response := PrintChainResponse{
		Blocks: blocks,
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
