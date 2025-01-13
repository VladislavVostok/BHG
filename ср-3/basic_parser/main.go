package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Status struct {
	ServerTime string
}

func main() {

	resp, err := http.Get("https://api.binance.us/api/v3/time")

	if err != nil {
		log.Panicln(err)
	}

	var status Status

	if resp.StatusCode == 200 {

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(string(body))

		if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
			log.Fatalln(err)
		}
	}
	defer resp.Body.Close()

	log.Printf("%s\n", status.ServerTime)
}
