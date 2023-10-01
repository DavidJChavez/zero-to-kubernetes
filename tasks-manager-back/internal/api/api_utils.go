package api_utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
)

// Common headers
const (
	HeaderContentType = "Content-Type"
)

// Common content types
const (
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
	ContentTypeHTML      = "text/html"
	ContentTypePlainText = "text/plain"
)

func WriteJson[T interface{}](w http.ResponseWriter, statusCode int, structData T) error {
	jsonData, err := json.Marshal(structData)
	if err != nil {
		return err
	}
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(statusCode)
	w.Write(jsonData)
	return nil
}

func WriteString(w http.ResponseWriter, statusCode int, s string) {
	w.Header().Set(HeaderContentType, ContentTypePlainText)
	w.WriteHeader(statusCode)
	w.Write([]byte(s))
}

func GetStructFromJson[T interface{}](r io.Reader) (T, error) {
	decoder := json.NewDecoder(r)
	var jsonData T
	err := decoder.Decode(&jsonData)
	if err != nil {
		return jsonData, err
	}
	st := reflect.TypeOf(jsonData)
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		log.Println(f.Tag.Get("required"))
	}

	return jsonData, nil
}
