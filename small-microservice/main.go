package main

import (
	"log"
)

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggingService(service)

	ApiServer := NewApiServer(service)
	log.Fatal(ApiServer.Start(":8080"))

	// fact, err := service.GetCatFact(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", fact)
}

//zeby to zbudowac i odpalic wystarczy "make run"