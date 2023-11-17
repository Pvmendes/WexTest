package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my-transaction-app/pkg/config"
	"my-transaction-app/pkg/transactions"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	service *transactions.Service
	trans   *transactions.Transaction
)

func main() {
	// load env variables just once in here so can be use in any other place
	config.InitEnvConfigs()

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Initialize services, repository, and api client
	service = transactions.NewService()

	// Setup HTTP routes for storing and retrieving transactions
	myRouter.HandleFunc("/store", storeTransactionHandler).Methods("POST")
	myRouter.HandleFunc("/retrieve/{id}", retrieveTransactionHandler).Methods("GET")
	myRouter.HandleFunc("/retrieve/{id}/{targetCountry}", retrieveTransactionHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// Define handlers for the routes
func storeTransactionHandler(rw http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	jsonDecoder := json.NewDecoder(r.Body)

	err := jsonDecoder.Decode(&trans)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	tId, err := service.StoreTransaction(trans)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	fmt.Fprint(rw, "Saved Transaction Id "+tId)
}

func retrieveTransactionHandler(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	keyId := vars["id"]
	keyTargetCountry := vars["targetCountry"]
	fmt.Println(keyTargetCountry)
	t, err := service.RetrieveTransaction(keyId, keyTargetCountry)

	if err != nil {
		fmt.Fprint(rw, err.Error())
	}

	json.NewEncoder(rw).Encode(t)
}
