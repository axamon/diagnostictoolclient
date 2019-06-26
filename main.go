package main

import (
        "fmt"
        "log"
        "net/http"
        "encoding/base64"
        "strings"
)

func main() {

        certPath := "server.pem"
        keyPath := "server.key"

//      api := NewAPI(certPath, keyPath)


        http.HandleFunc("/diagnostictool", func (w http.ResponseWriter, r *http.Request) {

                if basicAuth(w, r) {
                keys, _ := r.URL.Query()["tgu"]
                result, err := DiagnosticToolClient(keys[0])
                if err !=nil {
                        log.Println("funzione DiagnsticToolClient in errore: ", err.Error())

                }
                fmt.Fprint(w, result+"\n")
                return
                }
                w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
                w.WriteHeader(401)
                w.Write([]byte("401 Unauthorized\n"))
                return
        })


        err :=  http.ListenAndServeTLS(":8080", certPath, keyPath, nil)
//        err :=  http.ListenAndServe(":8080",nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}


func basicAuth(w http.ResponseWriter, r *http.Request) bool {

        auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

        if len(auth) != 2 || auth[0] != "Basic" {
                log.Println("Errore non Basic e non user e pass")
            http.Error(w, "authorization failed", http.StatusUnauthorized)
            return false
        }

        payload, _ := base64.StdEncoding.DecodeString(auth[1])
        pair := strings.SplitN(string(payload), ":", 2)

        if len(pair) != 2 || !validate(pair[0], pair[1]) {
                log.Println("Errore nella validazione user e pass")
            http.Error(w, "authorization failed", http.StatusUnauthorized)
            return false
        }

        return true
    }

func validate(username, password string) bool {
    if username == "test" && password == "test" {
        return true
    }
    return false
}
