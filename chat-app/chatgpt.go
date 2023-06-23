package main

import (
	"log"
	"net/http"

	"gopkg.in/olahol/melody.v1"
)

func main() {
	m := melody.New()

	// Handler function for handling incoming WebSocket connections
	m.HandleConnect(func(s *melody.Session) {
		log.Printf("Client connected: %s", s.Request.RemoteAddr)
	})

	// Handler function for handling incoming WebSocket messages
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Printf("Message received from client: %s", string(msg))

		// Broadcast the received message to all connected clients
		m.Broadcast(msg)
	})

	// Handler function for handling WebSocket disconnections
	m.HandleDisconnect(func(s *melody.Session) {
		log.Printf("Client disconnected: %s", s.Request.RemoteAddr)
	})

	// HTTP handler for upgrading the connection to WebSocket
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := m.HandleRequest(w, r)
		if err != nil {
			log.Printf("Failed to upgrade to WebSocket: %v", err)
		}
	})

	// Serve the HTML page that contains the WebSocket client code
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Start the server
	log.Println("WebSocket server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}


In this example:

The melody.New() function is used to create a new instance of the Melody server.

The m.HandleConnect function sets a handler for when a client connects to the WebSocket server.

The m.HandleMessage function sets a handler for receiving WebSocket messages from clients.

The m.HandleDisconnect function sets a handler for when a client disconnects from the WebSocket server.

The /ws endpoint is used to upgrade the HTTP connection to a WebSocket connection using m.HandleRequest.

The http.Handle function is used to serve the HTML page containing the WebSocket client code.

Finally, the server is started using http.ListenAndServe.

Make sure to install the Melody library by running go get gopkg.in/olahol/melody.v1 before running this code.

This is a basic example to get you started. You can customize and enhance it according to your specific requirements.