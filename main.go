package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {

        certPath := "server.pem"
        keyPath := "server.key"

//      api := NewAPI(certPath, keyPath)


        http.HandleFunc("/diagnostictool", func (w http.ResponseWriter, r *http.Request) {
                keys, _ := r.URL.Query()["tgu"]
                result, err := DiagnosticToolClient(keys[0])
                if err !=nil {
        }
                fmt.Fprint(w, result)
        })


        err :=  http.ListenAndServeTLS(":8080", certPath, keyPath, nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
