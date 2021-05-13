package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"-"` //invisivel no JSON
}

func main() {
	///////////////////// STRUCT EM JSON /////////////////////

	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroJSON)

	fmt.Println(bytes.NewBuffer(cachorroJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}
	cachorro2JSON, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(cachorro2JSON))

	///////////////////// JSON EM STRUCT /////////////////////
	cachorroEmJSON := `{"nome":"Thor","raca":"Pitbull","idade":9}`
	var ca cachorro

	err = json.Unmarshal([]byte(cachorroEmJSON), &ca)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ca)

}
