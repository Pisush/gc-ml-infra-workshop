package main

import (
	"fmt"
	"log"
	"net/http"

	aero "github.com/aerospike/aerospike-client-go"
)

const (
	// csv-realted
	namespace = "test"
	setName   = "creditcard"

	// BinMap field names
	userIDBin  = "UserID"
	setNameBin = "set_name"
	ccIDBin    = "ID"
	amountBin  = "AmountBin"
	classBin   = "ClassBin"
	timeBin    = "TimeBin"

	localhost         = "127.0.0.1"
	mlModelServingURL = "http://" + localhost + ":8501/v1/models/fraud:predict"
	inputLength       = 29
	fraudThreshold    = 0.5
)

var aeroClient *aero.Client

// webTxn is a struct for txn incoming in a web request
type webTxn struct {
	Timestamp string
	Amount    float64
	UserID    string
	CCNumHash string
	SellerID  int
	ItemID    int
}

// enrichedTxn is a struct used for sending to the ML model
type enrichedTxn struct {
}

// prediction is a struct for gettingt the prediction from the ML mode
type prediction struct {
}

// predictionHandler is the entry point to the system,
// ends in validating the prediction.
func predictionHandler(w http.ResponseWriter, req *http.Request) {
	// read txn, decode JSON, store in Aerospike
	incomingTxn := webTxn{}
	err := acceptTxn(req, aeroClient, &incomingTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// read txn by userID
	enrichedTxn := enrichedTxn{}
	txnOutcome, err := enrichTxn(aeroClient, &incomingTxn, &enrichedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send enriched txn to model serving web service
	modelPrediction, err := getPrediction(&enrichedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// compare prediction with classification
	validatePrediction(txnOutcome, modelPrediction)
}

// acceptTxn reads an incoming txn and stores it in the DB.
func acceptTxn(req *http.Request, client *aero.Client, incomingTxn *webTxn) (err error) {
	// read and decode txn

	// store the incoming txn in Aerospike

	return nil
}

// tryFloat tries to type assert v to float and return the value.
// Otherwise, it will return the default value dflt.
func tryFloat(v interface{}, dflt float64) float64 {
	f, ok := v.(float64)
	if !ok {
		return dflt
	}
	return f
}

// enrichTxn creates the enriched txn based on the given UserID.
func enrichTxn(aeroClient *aero.Client, incomingTxn *webTxn, enrichedTxn *enrichedTxn) (txnOutcome string, err error) {
	// read enriched data by UserID

	// marshal record.Bins in this struct and build
	// first build the inner array: [log(amount),v1,...v28]
	// using tryFloat()

	//next populate 2D array

	return "", err
}

// getPrediction sends the enriched txn to the model and gets a prediction.
func getPrediction(enrichedTxn *enrichedTxn) (modelPrediction string, err error) {
	// marshal to the expected struct

	// prepare request

	// make request

	// read response

	// unmarshal to response

	// check threshold to decide whether fraud

	return "", nil
}

// validatePrediction compares the model prediction
// with the classification from the DB.
func validatePrediction(txnOutcome, modelPrediction string) {
	// compare both predictions

	// advanced: run comparison for all fields in csv
	return
}

func main() {
	// set up a single instance of an Aerospike client
	// connection, it handles the connection pool internally
	var err error
	aeroClient, err = aero.NewClient(localhost, 3000)
	if err != nil {
		log.Fatal(err)
		return
	}

	// listen and serve
	fmt.Println("Listening on port :8090")
	http.HandleFunc("/", predictionHandler)
	http.ListenAndServe(":8090", nil)
}
