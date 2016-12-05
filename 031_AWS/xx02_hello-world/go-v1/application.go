package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
        if port == "" {
            port = "5000"
        }

        f, _ := os.Create("/var/log/golang/golang-server.log")
        defer f.Close()
        log.SetOutput(f)

        const indexPage = "public/index.html"
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            if r.Method == "POST" {
                if buf, err := ioutil.ReadAll(r.Body); err == nil {
                    log.Printf("Received message: %s\n", string(buf))
                }
            } else {
                log.Printf("Serving %s to %s...\n", indexPage, r.RemoteAddr)
                http.ServeFile(w, r, indexPage)
            }
        })

        http.HandleFunc("/scheduled", func(w http.ResponseWriter, r *http.Request){
            if r.Method == "POST" {
            log.Printf("Received task %s scheduled at %s\n", r.Header.Get("X-Aws-Sqsd-Taskname"), r.Header.Get("X-Aws-Sqsd-Scheduled-At"))
            }
        })

        log.Printf("Listening on port %s\n\n", port)
        http.ListenAndServe(":"+port, nil)
}
