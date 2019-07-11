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
	"github.com/tkanos/gonfig"
)

import (
	"strconv"
	"strings"
	"encoding/hex"
//	"github.com/shirou/gopsutil/host"
	"context"
	"log"
	"os"
	"fmt"
	"time"
)

// Configuration contiene gli elemnti per configurare il tool.
type Configuration struct 
{
	Token string `json:"token"`
}

var configuration Configuration

var scadenza int
var username, password string

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := gonfig.GetConf("conf.json", &configuration)
	if err != nil {
		log.Printf("ERROR Problema con il file di configurazione conf.json: %s\n", err.Error())
	}

	credBlob, _ := hex.DecodeString(configuration.Token)
	userEpass := string(decrypt(credBlob))
	credenziali := strings.Split(userEpass, " ")

	scadenza, err = strconv.Atoi(credenziali[0])
	if err != nil {
		log.Printf("ERROR Impossibile parsare scadenza del token: %s\n", err.Error())
	}
	username = credenziali[1]
	password = credenziali[2]



	oggi := time.Now().Unix()

	if oggi > int64(scadenza) {
		fmt.Println("Token scaduto. Impossibile proseguire.")
		os.Exit(1)
	}


	
	// var h []string
	//h, err := host.InfoWithContext(ctx)

	//fmt.Println(h.HostID)

	dati, err := diagnostictoolClient(ctx, os.Args[1])
	if err != nil {
		log.Printf("ERROR Impossiibile recuparare dati: %s", err.Error())
	}

	fmt.Println(dati)
}