package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/getAllFromVault", getAllFromVault).Methods("POST")
	myRouter.HandleFunc("/putToVault", putToVault).Methods("PUT")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}

func getAllFromVault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllFromVault")
	/*
		curl -X 'POST'  'https://vault.immudb.io/ics/api/v1/ledger/default/collection/default/documents/search' \
			-H 'accept: application/json' \
			-H 'X-API-Key: defaultro.QSrJ-6t3UOWiuvSe6DhTdg.rb1wj5FirC6fgbypVe-uqjtGM0CSCBehnGdu-Pw5ADwba8iY' \
			-H 'Content-Type: application/json' \
			-d '{"page":1,"perPage":100}'
	*/
	url := "https://vault.immudb.io/ics/api/v1/ledger/default/collection/default/documents/search"
	contentType := "application/json"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", contentType)
	//TODO: get from env
	req.Header.Add("X-API-Key", "defaultro.QSrJ-6t3UOWiuvSe6DhTdg.rb1wj5FirC6fgbypVe-uqjtGM0CSCBehnGdu-Pw5ADwba8iY")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
	var jsonResponse GetAllFromVaultResponse
	json.Unmarshal(body, &jsonResponse)
	var errorResponse ErrorResponse
	json.Unmarshal(body, &errorResponse)

	if errorResponse.Error != "" {
		err = json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func putToVault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: putToVault")
	/*
		curl -X 'PUT'   'https://vault.immudb.io/ics/api/v1/ledger/default/collection/default/document' \
			-H 'accept: application/json' \
			-H 'X-API-Key: default.QlRqXC-VqhIfjYVnuh-FRQ.W8qE8zUG4zQAm8hk4zomnyZysRe12_aCTmVUJa6Iit6-tk8Y' \
			-H 'Content-Type: application/json' \
			-d '{
				"accountNumber": "0000001",
				"accountName": "My first account",
				"iban": "XXX",
				"address": "123 Main Street",
				"amount": 123456789.1234,
				"type": "sending"
			}'
	*/
	url := "https://vault.immudb.io/ics/api/v1/ledger/default/collection/default/document"
	contentType := "application/json"

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", contentType)
	//TODO: get from env
	req.Header.Add("X-API-Key", "default.QlRqXC-VqhIfjYVnuh-FRQ.W8qE8zUG4zQAm8hk4zomnyZysRe12_aCTmVUJa6Iit6-tk8Y")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
	var jsonResponse PutToVaultResponse
	json.Unmarshal(body, &jsonResponse)

	if jsonResponse.DocumentId == "" {
		var errorResponse ErrorResponse
		json.Unmarshal(body, &errorResponse)
		err = json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Begin objects for JSON
type GetAllFromVaultResponse struct {
	/*
		{"page":1,"perPage":100,"revisions":[{"document":{"_id":"6567ad900000000000000004350e4972","_vault_md":{"creator":"a:c905cf65-cd16-4a95-a138-9a51b12ebd18","ts":1701293456},"address":"123 Main Street","age":30,"city":"New York","country":"USA","email":"johndoe@example.com","id":1,"is_active":true,"name":"John Doe","phone":"+1-123-456-7890","timestamp":"2023-05-10T12:00:00Z"},"revision":"","transactionId":""}],"searchId":""}
	*/
	Page      int    `json:"page"`
	PerPage   int    `json:"perPage"`
	SearchId  string `json:"searchId"`
	Revisions []struct {
		Document struct {
			DocumentId string `json:"_id"`
			VaultMeta  struct {
				Creator string `json:"creator"`
				Ts      int    `json:"ts"`
			} `json:"_vault_md"`
			AccountNumber string  `json:"accountNumber"`
			AccountName   string  `json:"accountName"`
			Iban          string  `json:"iban"`
			Address       string  `json:"address"`
			Amount        float64 `json:"amount"`
			Type          string  `json:"type"`
			Timestamp     string  `json:"timestamp"`
		} `json:"document"`
		Revision      string `json:"revision"`
		TransactionId string `json:"transactionId"`
	} `json:"revisions"`
}

type PutToVaultRequest struct {
	//structure: account number (unique), account name, iban, address, amount, type (sending, receiving)
	AccountNumber string  `json:"accountNumber"`
	AccountName   string  `json:"accountName"`
	Iban          string  `json:"iban"`
	Address       string  `json:"address"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"`
}

type PutToVaultResponse struct {
	//{"documentId":"6567b73b0000000000000005350e4973","transactionId":"6"}
	DocumentId    string `json:"documentId"`
	TransactionId string `json:"transactionId"`
}

type ErrorResponse struct {
	//{"code":409,"error":"unable to create document, error document already exists","status":"Conflict"}
	Code   int    `json:"code"`
	Error  string `json:"error"`
	Status string `json:"status"`
}
