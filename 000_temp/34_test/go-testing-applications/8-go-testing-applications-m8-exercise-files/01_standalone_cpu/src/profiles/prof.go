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

	f, err := os.Create("cpu.prof")

	if err != nil {
		log.Fatal(err.Error())
	}

	pprof.StartCPUProfile(f)
	for i := 0; i < 1000; i++ {
		http.Get("http://localhost:3000/api/purchaseOrders/1")
	}
	pprof.StopCPUProfile()

}
