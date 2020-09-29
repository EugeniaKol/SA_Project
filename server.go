package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)
func main(){
	http.HandleFunc("/time", mainPage)

	port := ":8795"
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)

	if err != nil{
		log.Fatal("listen and serve", err)
	}
}

type ResTime struct{
	TimeVal time.Time `json:"time"`
}

func mainPage(w http.ResponseWriter, r *http.Request){
	t := time.Now()
	t.Format(time.RFC3339)
	resTime := ResTime{t}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(resTime)

	//js, _ := json.Marshal(resTime)
	//w.Write(js)
}