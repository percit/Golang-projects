package  main

import (
	"fmt"
	"net/http"
)
func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world") //this just prints "hello world" on page
}


func main() {
	http.HandleFunc("/", helloWorldPage)//it associates path with handler function
	http.ListenAndServe(":8080", nil) //this starts the server
}
