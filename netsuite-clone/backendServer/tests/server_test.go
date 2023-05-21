package tests

import (
	"testing"
	"io/ioutil"
    "net/http"
    "net/http/httptest"
	"strings"
	"os"

	"backendServer/server"

	"github.com/gin-gonic/gin"
)

//running tests: "go test -v"
// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
  
//this runs all tests and exit program with status code after they are finished
	os.Exit(m.Run())
}


func TestPing(t *testing.T) {
    r := gin.Default()
    routes := server.Routes{}
    
    r.GET("/ping", routes.Ping)
  
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/ping", nil)
  
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
	
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "pong") > 0
	
		return statusOK && pageOK
	})
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()//recorder is sth to capture whole http response, aka status code, header and message body
  
	// Create the service and process the above request.
	r.ServeHTTP(w, req)//this processes request and writes response to recorder
  
	if !f(w) {
		t.Fail()
	}
}
  