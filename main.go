package main

import (
	"log"

	sw "receipt-inator/receipt-processor/go"
)

func main() {
	routes := sw.ApiHandleFunctions{}

	log.Println("Server started on port 8080")

	router := sw.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
