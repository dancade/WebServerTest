package main

import (
	"encoding/json"
	"os"
	"time"
	"net/http"
)

type HandsOn struct{
	Time 		time.Time `json:"time"`
	Hostname 	string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	resp := HandsOn{
		Time: time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	
	jsonResp, err := json.Marshal(&resp)
	if err != nil{
		w.Write([]byte("Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
//   this is a comment
func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":9090", nil)
}