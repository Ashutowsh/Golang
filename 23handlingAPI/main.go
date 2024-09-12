package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for building - file

type Building struct {
	BuildingId string `json : "buildingId"`
	Name       string `json : "Name of the building"`
	Floors     []*Floor
}

type Floor struct {
	FloorId string `json : "floorId"` // `json : "-"` --> Will not add this data into JSON.
	FloorNo string `json : "Floor number"`
	Rooms   []*Room
}

type Room struct {
	RoomId   string `json : "roomId"`
	AC_NonAC string `json : "AC/Non-AC"`
	RoomNo   string `json : "Room Number"`
}

// fake DB
var buildings []Building

func (b *Building) IsEmpty() bool {
	// return b.BuildingId == "" && b.Name == ""
	return b.Name == ""
}

func (b *Building) numberOfFloors() int {
	return len(b.Floors)
}

func (b *Building) numberOf() int {
	return len(b.Floors)
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	buildings = append(buildings, Building{
		BuildingId: "B001",
		Name:       "Main Building",
		Floors: []*Floor{
			{
				FloorId: "F001",
				FloorNo: "1",
				Rooms: []*Room{
					{
						RoomId:   "R001",
						AC_NonAC: "AC",
						RoomNo:   "101",
					},
					{
						RoomId:   "R002",
						AC_NonAC: "Non-AC",
						RoomNo:   "102",
					},
				},
			},
			{
				FloorId: "F002",
				FloorNo: "2",
				Rooms: []*Room{
					{
						RoomId:   "R003",
						AC_NonAC: "AC",
						RoomNo:   "201",
					},
					{
						RoomId:   "R004",
						AC_NonAC: "Non-AC",
						RoomNo:   "202",
					},
				},
			},
		},
	},
		Building{
			BuildingId: "B002",
			Name:       "Annex Building",
			Floors: []*Floor{
				{
					FloorId: "F003",
					FloorNo: "1",
					Rooms: []*Room{
						{
							RoomId:   "R005",
							AC_NonAC: "AC",
							RoomNo:   "101",
						},
						{
							RoomId:   "R006",
							AC_NonAC: "Non-AC",
							RoomNo:   "102",
						},
					},
				},
				{
					FloorId: "F004",
					FloorNo: "2",
					Rooms: []*Room{
						{
							RoomId:   "R007",
							AC_NonAC: "AC",
							RoomNo:   "201",
						},
						{
							RoomId:   "R008",
							AC_NonAC: "Non-AC",
							RoomNo:   "202",
						},
					},
				},
			},
		})

	r.HandleFunc("/building/{id}", getOneBuilding).Methods("GET")
	r.HandleFunc("/buildings", getAllBuildings).Methods("GET")
	r.HandleFunc("/building", createBuilding).Methods("POST")
	r.HandleFunc("/building/{id}", updateOneBuilding).Methods("PUT")
	r.HandleFunc("/building/{id}", deleteOneBuilding).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Builder.com</h1>"))

}

func getAllBuildings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Buildings.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buildings)
}

func getOneBuilding(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Buildings.")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, building := range buildings {
		if building.BuildingId == params["id"] {
			json.NewEncoder(w).Encode(building)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with id = " + params["id"])
}

func createBuilding(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Building.")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var building Building
	_ = json.NewDecoder(r.Body).Decode(&building)

	if building.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	building.BuildingId = strconv.Itoa(rand.Intn(100))
	buildings = append(buildings, building)
	json.NewEncoder(w).Encode(building)
}

func updateOneBuilding(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Building.")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	params := mux.Vars(r)

	for index, building := range buildings {
		if building.BuildingId == params["id"] {
			buildings = append(buildings[:index], buildings[index+1:]...)
			var building Building
			_ = json.NewDecoder(r.Body).Decode(&building)

			if building.IsEmpty() {
				json.NewEncoder(w).Encode("No data inside JSON")
				return
			}

			building.BuildingId = params["id"]
			buildings = append(buildings, building)

			json.NewEncoder(w).Encode(building)
			return // Not required as we send the response it is automatically returned.
		}
	}
}

func deleteOneBuilding(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course.")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	params := mux.Vars(r)

	for index, building := range buildings {
		if building.BuildingId == params["id"] {
			buildings = append(buildings[:index], buildings[index+1:]...)
			break
		}
	}
}
