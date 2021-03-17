package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

/*

This is never used, so we'll comment it out.

type PageVariables struct {
	Date string
	Time string
}

*/

type Item struct {
	ID          string `json:"ID"`
	Item        string `json:"Item"`
	Description string `json:"Description"`
	Price       string `json:"Price"`
}

type allItems []Item

var Items = allItems{
	{
		ID:          "1",
		Item:        "Iphone",
		Description: "bla bla bla",
		Price:       "1000$",
	},
	{
		ID:          "3",
		Item:        "Drip",
		Description: "To be cool",
		Price:       "20000$",
	},
	{
		ID:          "4",
		Item:        "Laptop",
		Description: "Normal Laptop",
		Price:       "300$",
	},
	{
		ID:          "5",
		Item:        "Pc gaming",
		Description: "powerful pc",
		Price:       "2000$",
	},
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there read the docs")
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item
	reqBody, err := ioutil.ReadAll(r.Body) // Convert r.Body into a readable format
	if err != nil {
		fmt.Fprintf(w, "BRUH enter data with the event id, title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newItem)
	Items = append(Items, newItem)     // Add the newly created event to the array of events
	w.WriteHeader(http.StatusCreated)  // Return the 201 created status code
	json.NewEncoder(w).Encode(newItem) // Return the newly created event
}

func GetOneItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get the "id" URL query
	ItemID := mux.Vars(r)["id"]
	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleItem := range Items {
		if singleItem.ID == ItemID {
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Items)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get the "id" URL query
	ItemID := mux.Vars(r)["id"]
	var updatedItem Item
	reqBody, err := ioutil.ReadAll(r.Body) // Convert r.Body into a readable format
	if err != nil {
		fmt.Fprintf(w, "BRUH enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedItem)
	for i, singleItem := range Items {
		if singleItem.ID == ItemID {
			singleItem.Item = updatedItem.Item
			singleItem.Description = updatedItem.Description
			Items[i] = singleItem
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get the "id" URL query
	ItemID := mux.Vars(r)["id"]
	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for i, singleItem := range Items {
		if singleItem.ID == ItemID {
			Items = append(Items[:i], Items[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", ItemID)
		}
	}
}
