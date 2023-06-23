Client struct functions:

type Client struct {
	username string
	color string
	// connection Connection
}
 

connect(color string)
- makes websocket connection to the server specified by url argument
- assigns a color to color field of struct (to show in GUI)
- assigns a connection to the connection field of Client

send(message string) 
- sends a message to server using websocket connection 
- takes provided message, formats it to JSON with "username", "color", "message"
- (maybe) the json is encrypted
- encrypted message is sent to server as a websocket text message

recieveHandler()
- this is a goroutine
- handles incoming data from server and unmarchals recieved JSON into a structure.Message struct
- reads all data from server, decrypts it and unmarshals JSON
- the data is shown somewhere