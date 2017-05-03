package main

import (
	"encoding/json"
	"github.com/clearfunction/go-lunch/human"
	"log"
)

func main() {
	var daniel human.Human
	var jsonBlob = []byte(`
	  { "personName": "Daniel", "personAge": 35 }
	`)

	err := json.Unmarshal(jsonBlob, &daniel)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", daniel)
}
