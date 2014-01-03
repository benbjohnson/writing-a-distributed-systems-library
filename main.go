package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/goraft/raft"
)

var server raft.Server
var currentValue int

func init() {
	raft.RegisterCommand(&AddCommand{})
}

func main() {
	hostname, _ := os.Hostname()

	// Initialize the internal Raft server.
	transporter := raft.NewHTTPTransporter("/raft")
	server, _ = raft.NewServer(hostname, ".", transporter, nil, nil, "")

	// Attach the Raft server to the HTTP server.
	transporter.Install(server, http.DefaultServeMux)

	// Create a /add endpoint.
	http.HandleFunc("/add", addHandler)

	// Start both servers.
	server.Start()

	// Initialize to a cluster of one for this example.
	if server.IsLogEmpty() {
		_, err := server.Do(&raft.DefaultJoinCommand{Name:hostname})
		if err != nil {
			log.Fatal(err)
		}
	}

	// Initialize HTTP server.
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// addHandler executes a command against the raft server and returns the result.
func addHandler(w http.ResponseWriter, req *http.Request) {
	value, _ := strconv.Atoi(req.FormValue("value"))
	newValue, err := server.Do(&AddCommand{Value: value})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%d\n", newValue)
}


// AddCommand adds a number to the current value of the system.
type AddCommand struct {
	Value int
}

func (c *AddCommand) CommandName() string {
	return "add"
}

func (c *AddCommand) Apply(ctx raft.Context) (interface{}, error) {
	currentValue += c.Value
	return currentValue, nil
}
