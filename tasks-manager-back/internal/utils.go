package internal

import (
	"encoding/json"
	"log"
)

func PrintStruct[T interface{}](a T) {
	jsonA, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonA))
}
