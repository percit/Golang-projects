package main

import (
	"gopherconuk/homepage"
	"gopherconuk/server"
	"log"
	"net/http"
	"os"
)

var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)//to jest tworzenie specyficznego loggera


	h := homepage.NewHandlers(logger) //to tak naprawde jest dependency injection(wrzucasz wskaznik do konstruktora klasy)

	mux := http.NewServeMux() //on odpowiada za http itd
	h.SetupRoutes(mux)

	srv := server.New(mux, GcukServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
