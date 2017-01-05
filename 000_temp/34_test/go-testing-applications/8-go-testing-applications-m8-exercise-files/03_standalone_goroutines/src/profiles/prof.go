package main

import (
	"log"
	"net/http"
	"os"
	"poms"
	"poms/ctrl"
	"runtime/pprof"
)

func main() {
	ctrl.Setup()

	go http.ListenAndServe(":3000", new(poms.GZipServer))

	f, err := os.Create("goroutines.prof")

	if err != nil {
		log.Fatal(err.Error())
	}
	for i := 0; i < 5; i++ {
		http.Get("http://localhost:3000/api/purchaseOrders/1")
	}

	pprof.Lookup("goroutine").WriteTo(f, 1)
}
