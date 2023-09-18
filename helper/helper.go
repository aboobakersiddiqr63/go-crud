package helper

import (
	"fmt"
	"log"
	"net/http"
)

func SetCommonHeaders(w http.ResponseWriter, method string) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", method)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func GetCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func HandleException(err error, funcName string) {
	if err != nil {
		fmt.Printf("Error in %v\n", funcName)
		log.Fatal(err)
	}
}
