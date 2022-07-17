package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Bike struct {
	Id     string `json:"id"`
	Type   string `json:"title"`
	Year   string `json:"year"`
	Colour string `json:"colour"`
	Owner  *Owner `json:"owner"`
}

type Owner struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

var bikes []Bike

func getBikes(w http.ResponseWriter, r *http.Request) {
	// response content in json
	w.Header().Set("Content-Type", "application/json")
	// encoding bikes slice
	json.NewEncoder(w).Encode(bikes)
}

func getBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for _, item := range bikes {
		if item.Id == parameters["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getBikeByYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	var yearBikes []Bike
	for _, item := range bikes {
		if item.Id == parameters["id"] {
			yearBikes = append(yearBikes, item)
		}
	}
	json.NewEncoder(w).Encode(yearBikes)
}

func getBikeByPhone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	var phoneBikes []Bike
	for _, item := range bikes {
		if item.Owner.Phone == parameters["phone"] {
			phoneBikes = append(phoneBikes, item)
		}
	}
	json.NewEncoder(w).Encode(phoneBikes)
}

func deleteBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// index to delete as a part of request
	parameters := mux.Vars(r)
	for index, item := range bikes {
		if item.Id == parameters["id"] {
			// delete item with actual index
			bikes = append(bikes[:index], bikes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bikes)
}

func createBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// create variable bike of type Bike
	var bike Bike
	// decode requested bike type Bike
	_ = json.NewDecoder(r.Body).Decode(&bike)
	// generate random index
	bike.Id = strconv.Itoa(rand.Intn(2000000))
	bikes = append(bikes, bike)
	json.NewEncoder(w).Encode(bike)

}

func updateBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range bikes {
		if item.Id == parameters["id"] {
			// delete item with actual index
			bikes = append(bikes[:index], bikes[index+1:]...)
			var bike Bike
			_ = json.NewDecoder(r.Body).Decode(&bike)
			bike.Id = parameters["id"]
			bikes = append(bikes, bike)
			json.NewEncoder(w).Encode(bike)

		}
	}
}

func main() {
	r := mux.NewRouter()

	bikes = append(bikes, Bike{Id: "1", Type: "Cross", Year: "2017", Colour: "yellow", Owner: &Owner{FirstName: "Gosia", LastName: "Dutka", Phone: "-", Email: "gosia@gmail.com"}})

	r.HandleFunc("/bikes", getBikes).Method("GET")
	r.HandleFunc("/bikes/{id}", getBike).Method("GET")
	r.HandleFunc("/bikes/{year}", getBikeByYear).Method("GET")
	r.HandleFunc("/bikes/{year}", getBikeByPhone).Method("GET")
	r.HandleFunc("/bikes", createBike).Method("CREATE")
	r.HandleFunc("/bikes/{id}", updateBike).Method("UPDATE")
	r.HandleFunc("/bikes/{id}", deleteBike).Method("DELETE")

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
