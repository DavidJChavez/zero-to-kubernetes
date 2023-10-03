package pkg

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintStruct[T interface{}](a T) {
	jsonA, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonA))
}

func PrintErr(v any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, v)
	log.Println(colored)
}

func PrintInfo(v any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, v)
	log.Println(colored)
}
