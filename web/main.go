package  main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func helloWorldPageTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world") //this just prints "hello world" on page
}

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			fmt.Fprint(w, "hello world")
		case "/ninja":                       	
			fmt.Fprint(w, "Marcin")
		default:
			fmt.Fprint(w, "Error")
	}
}

func htmlHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //this specifies, that we are outputting html code
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func main() {
	//http.HandleFunc("/", helloWorldPageTest)//it associates path with handler function
	//http.HandleFunc("/", helloWorldPage)//this prints with different texts and paths
	//http.HandleFunc("/", htmlHelloWorld)
	//http.ListenAndServe(":8080", nil) //this starts the server


	response, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
	}
	
}
