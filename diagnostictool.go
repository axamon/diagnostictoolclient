package main

import (
        "fmt"
        "encoding/json"
        "regexp"
        "crypto/tls"
        "io/ioutil"
        "log"
        "net/http"
        "os"
)

const (
        endpointTgu = "https://10.38.34.138:8443/DiagnosticTool/api.php?method=DiagnosticTool&sincrono=N&format=json&tgu="
        endpointEsito = "https://10.38.34.138:8443/DiagnosticTool/api.php?method=DiagnosticTool&sincrono=Y&format=json&cod_esito="
)

type response struct {
Esito string  `json:"esito"`
TDResponseCod string  `json:"responsecode"`
TDResponse string  `json:"response"`
CodEsito string   `json:"cod_esito"`
}



var re = regexp.MustCompile(`(?m)^\d+$`)

func DiagnosticToolClient(tgu string) (result string, err error){

        fmt.Println(tgu)

        // Costringe il client ad accettare anche certificati https non validi
        // o scaduti.
        transCfg := &http.Transport{
                // Ignora certificati SSL scaduti.
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        client := &http.Client{Transport: transCfg}
        url := endpointTgu + tgu

        fmt.Println(url)

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                log.Printf("ERROR Impossibile creare richiesta: %s\n", err.Error())
        }


        username := os.Getenv("DiagnosticToolUsername")
        password := os.Getenv("DiagnosticToolPassoword")

        req.SetBasicAuth(username, password)

        resp, err := client.Do(req)
        if err != nil {
                log.Printf("ERROR Impossibile inviare richiesta http: %s\n", err.Error())
        }
        //defer resp.Body.Close()

        responsBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Printf("ERROR Impossibile leggere body reqest: %s\n", err.Error())
        }


                fmt.Println(string(responsBody))


                risposta := new(response)


                err = json.Unmarshal(responsBody, &risposta)
                if err != nil {
                        log.Println(err)
                }

                fmt.Println(risposta)
      

        dati, err :=  dt(risposta.CodEsito)
            if err != nil {
                        log.Println(err)
                }

        return dati, err

}



func dt(code string) (str string, err error){


        // Costringe il client ad accettare anche certificati https non validi
        // o scaduti.
        transCfg := &http.Transport{
                        // Ignora certificati SSL scaduti.
                        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        client := &http.Client{Transport: transCfg}
        url := endpointEsito + code

        fmt.Println(url)

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                        log.Printf("ERROR Impossibile creare richiesta: %s\n", err.Error())
        }


        username := os.Getenv("DiagnosticToolUsername")
        password := os.Getenv("DiagnosticToolPassoword")

        req.SetBasicAuth(username, password)

        resp, err := client.Do(req)
        if err != nil {
                        log.Printf("ERROR Impossibile inviare richiesta http: %s\n", err.Error())
        }
        //defer resp.Body.Close()

        responsBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                        log.Printf("ERROR Impossibile leggere body reqest: %s\n", err.Error())
        }

        return string(responsBody), err
}
