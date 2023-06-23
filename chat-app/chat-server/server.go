package main

//VARIABLES
//clients variable thath will keep track of all connected WebSocket clients
//broadcast - channel for broadcasting messages to all connected clients
//upgrader - gorrila has this websocket.Upgrader for upgrading http connections to websocket connections RECHECK

import (
	"net/http"

	"gopkg.in/olahol/melody.v1"
)

func main() {


	// http.HandleFunc("/ws", handleConnections) //cos takiego do ogarniania polaczen
	// go handleMessages() 
	//podlaczenie sie pod port 8080
}


//FUNCTIONS:
//handleConnections - it will read connections and upgrade http to websocket in infinite loop, if there is an error, stop program

//handleMessages - a goroutine that will broadcast messages to all connected clients (aka prints them it); does this:
// - waits for message to be recieved from broadcast channel
// - iterates over all connected clients and writes messages to each client using client.WriteMessage()
// - if there is an error writing a message to client, remove the client


