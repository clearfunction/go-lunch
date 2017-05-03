package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// type hints allow for JSON parsing of mismatched names
type Human struct {
	Name string `json:"personName"`
	Age  int    `json:"personAge"`
}

func main() {
	var daniel Human

	var jsonBlob = []byte(`
	  { "personName": "Daniel", "personAge": 35 }
	`)

	// &daniel is address of identifier daniel; writes directly to existing object
	err := json.Unmarshal(jsonBlob, &daniel)

	// die on error with log.Fatal
	if err != nil {
		log.Fatal(err)
	}

	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v", daniel)
}
