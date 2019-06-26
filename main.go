// Copyright (c) 2019 Alberto Bregliano
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var port = flag.String("p", ":8080", "porta da usare, default :8080")

func main() {

	flag.Parse()

	// Crea il contesto
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Certificati per https
	// da creare con i comandi:
	//  openssl genrsa -out server.key 2048
	//  openssl ecparam -genkey -name secp384r1 -out server.key
	//  openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
	certPath := "server.pem"
	keyPath := "server.key"

	http.HandleFunc("/diagnostictool", func(w http.ResponseWriter, r *http.Request) {
		r.WithContext(ctx)

		switch ok := basicAuth(w, r); ok {
		// Se correttamente autenticato
		case true:
			keys, _ := r.URL.Query()["tgu"]

			// result contiene il json con i dati.
			result, err := DiagnosticToolClient(ctx, keys[0])
			if err != nil {
				log.Println("funzione DiagnsticToolClient in errore: ", err.Error())
			}

			// invia il json con i dati in risposta.
			fmt.Fprint(w, result+"\n")
			return

		// Se NON correttamente autenticato
		case false:
			w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}
	})

	err := http.ListenAndServeTLS(*port, certPath, keyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// basicAuth gestisce l'autenticazione all'app.
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
