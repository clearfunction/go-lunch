package main

import (
	"encoding/json"
	"github.com/clearfunction/go-lunch/human" // HL
	"log"
)

func main() {
	var daniel human.Human
	var jsonBlob = []byte(`
	  { "personName": "Daniel", "personAge": 35 }  // HL
	`)

	err := json.Unmarshal(jsonBlob, &daniel) // HL

	if err != nil {
		log.Fatal(err) // HL
	}

	log.Printf("%+v", daniel)
}
