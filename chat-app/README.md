Main plan:

1 sending messages with a sender's nick
- opening a websocket connection
2 showing all messages as printing





To do:
- logging in / making account
- sending messages
- recieving messages
- differentiating messages from different people
- Frontend


technologies I want to use:
- Go
- Websockets (as they are real time) -> Melody library
- Some kind of maintaining messages (some database on server side maybe)
- React (which I need to learn)


z tamtego projektu:
- ansible nie uzyjemy, to jest docker itd
- public to tez nie moje, bo to jest html + css, ja chce reacta
- w main.js sie wiekszosc dzieje, uzywa on websocketa i url golangowego
- przy kliknieciu jest wysylany text na websocket golangowy, a potem wypisywany nizej
ws.onmessage = function (msg) {
	console.log(msg.data)
    insertMessage(JSON.parse(msg.data))
};
- ten kod wyzej odpowiada za odczytywanie danych z websocketa



architecture:
- The server creates an instance of the Client type for each websocket connection.

- WebSocket Server: Set up a WebSocket server that will handle incoming client connections, manage WebSocket sessions, and facilitate communication between clients.
- User Management: Implement user authentication and user management functionality on the server-side, such as handling user registration, login, and storing user information.
- Message Handling: Develop logic for handling incoming messages from clients, including broadcasting messages to other connected clients, storing messages in a database, and retrieving message history.
- Event Handling: Implement event handlers on the server-side to trigger actions based on specific events, such as user joining or leaving the chat, and update the client interfaces accordingly.
- Persistence: If desired, integrate a database to store user information, chat history, and other relevant data.
- Add features such as private messaging, file sharing, emojis, or other chat-related functionalities to enhance the user experience.
- Implement additional security measures, such as encryption, to protect sensitive user information and messages.








przyklady:
https://github.com/olahol/melody tu jest libka do websocketow i przyklad takiego chatu
https://github.com/roblaszczak/simple-go-chat (to jest na jquery itd, wiec idk)
