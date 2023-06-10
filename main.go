package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
)

var sf *sonyflake.Sonyflake

func init() {
	log.Println("JALAN")

	var st sonyflake.Settings
	// st.MachineID = awsutil.AmazonEC2MachineID

	machineId, err := awsutil.AmazonEC2MachineID()
	log.Printf("%+v\n", machineId)
	log.Printf("%+v\n", err)

	log.Println("LEWAT")

	log.Printf("%+v\n", st)
	log.Printf("%+v\n", st.MachineID)
	sf = sonyflake.NewSonyflake(st)
	log.Println("LEWAT")
	log.Printf("%+v\n", sf)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := sf.NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(sonyflake.Decompose(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	w.Write(body)
}

func main() {
	log.Println("Hosted on 8080")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
