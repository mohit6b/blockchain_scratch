package router

import (
	"github.com/gorilla/mux"
	"github.com/mohit6b/pow-blockchain-go/src/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	print(router)
	router.HandleFunc("/createblockchain", handlers.HandleCreateBlockchain).Methods("POST")
	router.HandleFunc("/createwallet", handlers.HandleCreateWallet).Methods("POST")
	router.HandleFunc("/printchain", handlers.HandlePrintChain).Methods("POST")
	router.HandleFunc("/startnode", handlers.HandleStartNode).Methods("POST")
	router.HandleFunc("/getbalance", handlers.HandleGetBalance).Methods("POST")
	router.HandleFunc("/listaddress", handlers.HandleListAddress).Methods("POST")
	router.HandleFunc("/reindexutxo", handlers.HandleReindexUTXO).Methods("POST")
	router.HandleFunc("/send", handlers.HandleSend).Methods("POST")

	return router
}
