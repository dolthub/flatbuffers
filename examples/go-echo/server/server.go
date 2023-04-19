package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/flatbuffers/examples/go-echo/net"
)

func echo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Unable to read request body: %v\n", err)
		return
	}

	req := net.GetRootAsRequest(body, 0)
	player := req.Player(nil)

	fmt.Printf("Got request (name: %v, hp: %v)\n", string(player.Name()), player.Hp())
	w.Write(body)
}

func main() {
	http.HandleFunc("/echo", echo)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
